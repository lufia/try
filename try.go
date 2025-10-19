// Package try provides error-handling utilities.
package try

func getbp(skip int) uintptr

// Scope represents the fallback point.
type Scope struct {
	sp    uintptr
	bp    uintptr
	dx    uintptr
	pc    uintptr
	probe uintptr // BP of Handle's parent
	err   error
}

func waserror(s *Scope) bool
func raise(s *Scope) bool

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
	if err == nil {
		return
	}
	s.err = err

	bp := getbp(skip + 1)
	d := bp - s.probe
	ns := *s
	ns.sp += d
	ns.bp += d
	ns.dx += d
	raise(&ns)
	panic("do not reach here")
}

type Cond[T any] struct {
	v   T
	err error
	fn  func(err error) error
}

func (c *Cond[T]) Eval(s *Scope) T {
	err := c.err
	if err != nil {
		if c.fn != nil {
			err = c.fn(err)
		}
		s.raise(2, err)
	}
	return c.v
}

func (c *Cond[T]) Wrap(f func(err error) error) *Cond[T] {
	c.fn = f
	return c
}

type Cond2[T1, T2 any] struct {
	v1  T1
	v2  T2
	err error
	fn  func(err error) error
}

func (c *Cond2[T1, T2]) Eval(s *Scope) (T1, T2) {
	err := c.err
	if c.fn != nil && err != nil {
		err = c.fn(err)
	}
	s.raise(2, err)
	return c.v1, c.v2
}

func (c *Cond2[T1, T2]) Wrap(f func(err error) error) *Cond2[T1, T2] {
	c.fn = f
	return c
}

// Check checks whether err is not nil.
// If err is nil, it returns v.
// Otherwise it rewinds to the fallback point s, then [Handle] returns err.
//
// Check should be called on the same stack to [Handle].
func Check[T any](v T, err error) *Cond[T] {
	return &Cond[T]{v: v, err: err}
}

// Check2 is a variant of [Check].
func Check2[T1, T2 any](v1 T1, v2 T2, err error) *Cond2[T1, T2] {
	return &Cond2[T1, T2]{v1: v1, v2: v2, err: err}
}
