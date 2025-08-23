// Package try provides error-handling utilities.
package try

type Scope struct {
	pc  uintptr
	sp  uintptr
	err error
}

func waserror(s *Scope) bool
func raise(s *Scope)

// Handle creates a fallback point.
func Handle() (*Scope, error) {
	var s Scope
	if waserror(&s) {
		return &s, s.err
	}
	return &s, nil
}

// Raise fallbacks to the fallback point s.
func (s *Scope) Raise(err error) {
	if err == nil {
		return
	}
	s.err = err
	raise(s)
	panic("do not reach")
}

// Check checks whether err is not nil, if an error was happen, it jumps [Handle].
//
//go:noinline
func Check[T any](v T, err error) func(s *Scope) T {
	return func(s *Scope) T {
		s.Raise(err)
		return v
	}
}

// Check2 is a variant of [Check].
//
//go:noinline
func Check2[T1, T2 any](v1 T1, v2 T2, err error) func(s *Scope) (T1, T2) {
	return func(s *Scope) (T1, T2) {
		s.Raise(err)
		return v1, v2
	}
}
