// Package try provides error-handling utilities.
package try

// Scope represents the fallback point.
type Scope struct {
	sp    uintptr
	bp    uintptr
	ctxt  uintptr
	pc    uintptr
	probe uintptr // BP of Handle's parent

	err     error
	handler func(err error) error
}

// Option configures [Check], [Check1] and [Check2].
type Option func(*Scope)

func applyOpts(s *Scope, opts ...Option) {
	for _, opt := range opts {
		opt(s)
	}
}

func WithHandler(f func(err error) error) Option {
	return func(s *Scope) {
		s.handler = f
	}
}

func waserror(s *Scope) bool
func raise(s *Scope) bool
func getbp(skip int) uintptr

// Handle creates a fallback point.
func Handle() (*Scope, error) {
	var s Scope
	if waserror(&s) {
		return nil, s.err
	}
	return &s, nil
}

func (s *Scope) raise(skip int, err error) {
	if s.handler != nil {
		err = s.handler(err)
	}
	if err == nil {
		return
	}
	s.err = err

	bp := getbp(skip + 1)
	d := bp - s.probe
	s.probe += d
	s.sp += d
	s.bp += d
	s.ctxt += d
	raise(s)
	panic("do not reach here")
}

type Rewinder func(*Scope, ...Option)

// Check checks whether err is not nil.
// If err is nil, it does nothing.
// Otherwise it rewinds to the fallback point s, then [Handle] returns err.
//
// Check should be called on the same stack to [Handle].
func Check(err error) Rewinder {
	return func(s *Scope, opts ...Option) {
		applyOpts(s, opts...)
		s.raise(1, err)
	}
}

type Rewinder1[T any] func(*Scope, ...Option) T

// Check1 checks whether err is not nil.
// If err is nil, it returns v.
// Otherwise it rewinds to the fallback point s, then [Handle] returns err.
//
// Check1 should be called on the same stack to [Handle].
func Check1[T any](v T, err error) Rewinder1[T] {
	return func(s *Scope, opts ...Option) T {
		applyOpts(s, opts...)
		s.raise(1, err)
		return v
	}
}

type Rewinder2[T1, T2 any] func(*Scope, ...Option) (T1, T2)

// Check2 is a variant of [Check1].
func Check2[T1, T2 any](v1 T1, v2 T2, err error) Rewinder2[T1, T2] {
	return func(s *Scope, opts ...Option) (T1, T2) {
		applyOpts(s, opts...)
		s.raise(1, err)
		return v1, v2
	}
}
