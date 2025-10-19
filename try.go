// Package try provides error-handling utilities.
package try

// Scope represents the fallback point.
type Scope struct {
	sp    uintptr
	bp    uintptr
	dx    uintptr
	pc    uintptr
	probe uintptr // BP of Handle's parent

	err     error
	handler func(err error) error
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

// Raise fallbacks to the fallback point s.
//
// Raise should be called on the same stack to [Handle].
func (s *Scope) Raise(err error) {
	s.raise(1, err)
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
	s.dx += d
	raise(s)
	panic("do not reach here")
}

// Option configures [Check] and [Check2].
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

// Check checks whether err is not nil.
// If err is nil, it returns v.
// Otherwise it rewinds to the fallback point s, then [Handle] returns err.
//
// Check should be called on the same stack to [Handle].
func Check[T any](v T, err error) func(*Scope, ...Option) T {
	return func(s *Scope, opts ...Option) T {
		applyOpts(s, opts...)
		s.raise(1, err)
		return v
	}
}

// Check2 is a variant of [Check].
func Check2[T1, T2 any](v1 T1, v2 T2, err error) func(*Scope, ...Option) (T1, T2) {
	return func(s *Scope, opts ...Option) (T1, T2) {
		applyOpts(s, opts...)
		s.raise(1, err)
		return v1, v2
	}
}
