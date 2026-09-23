// Package try provides error-handling utilities.
package try

import (
	"errors"
	"fmt"
	"unsafe"
)

// Checkpoint represents the fallback point.
type Checkpoint[E error] struct {
	regs
	err      E
	handlers []func(err E) E
}

type regs struct {
	sp    uintptr
	bp    uintptr
	ctxt  uintptr
	pc    uintptr
	probe uintptr // BP of Handle's parent
}

// Option configures [Check], [Check1] and [Check2].
type Option[E error] func(*Checkpoint[E])

func applyOpts[E error](cp *Checkpoint[E], opts ...Option[E]) {
	for _, opt := range opts {
		opt(cp)
	}
}

func WithHandler[E error](f func(err E) E) Option[E] {
	return func(cp *Checkpoint[E]) {
		cp.handlers = append(cp.handlers, f)
	}
}

func WithDescription(format string, args ...any) Option[error] {
	prefix := fmt.Sprintf(format, args...)
	return func(cp *Checkpoint[error]) {
		cp.handlers = append(cp.handlers, func(err error) error {
			return fmt.Errorf("%s: %w", prefix, err)
		})
	}
}

func WithIgnore(errs ...error) Option[error] {
	return func(cp *Checkpoint[error]) {
		cp.handlers = append(cp.handlers, func(err error) error {
			for _, e := range errs {
				if errors.Is(err, e) {
					return nil
				}
			}
			return err
		})
	}
}

func waserror(cp uintptr) bool
func raise(cp uintptr) bool
func stkhi() uintptr

// Handle creates a fallback point.
func Handle() (*Checkpoint[error], error) {
	var cp Checkpoint[error]
	if waserror(uintptr(unsafe.Pointer(&cp))) {
		return nil, cp.err
	}
	return &cp, nil
}

// HandleFor creates a fallback point for the type argument E.
func HandleFor[E error]() (*Checkpoint[E], E) {
	var cp Checkpoint[E]
	if waserror(uintptr(unsafe.Pointer(&cp))) {
		return nil, cp.err
	}
	var zero E
	return &cp, zero
}

func isZero[T any](v T) bool {
	n := unsafe.Sizeof(v)
	if n == 0 {
		return true
	}
	b := unsafe.Slice((*byte)(unsafe.Pointer(&v)), n)
	for _, n := range b {
		if n != 0 {
			return false
		}
	}
	return true
}

func (cp *Checkpoint[E]) raise(skip int, err E) {
	if isZero(err) {
		return
	}
	for _, f := range cp.handlers {
		err = f(err)
		if isZero(err) {
			return
		}
	}
	cp.err = err

	hi := stkhi()
	d := hi - cp.probe
	cp.probe += d
	cp.sp += d
	cp.bp += d
	cp.ctxt += d
	raise(uintptr(unsafe.Pointer(cp)))
	panic("do not reach here")
}

// Rewind rewinds current execution point to cp.
func (cp *Checkpoint[E]) Rewind(err E) {
	cp.raise(1, err)
}

type RewinderFunc[E error] func(*Checkpoint[E], ...Option[E])

// Check checks whether err is not nil.
// If err is nil, it does nothing.
// Otherwise it rewinds to the fallback point s, then [Handle] returns err.
//
// Check should be called on the same stack to [Handle].
func Check[E error](err E) RewinderFunc[E] {
	return func(cp *Checkpoint[E], opts ...Option[E]) {
		applyOpts(cp, opts...)
		cp.raise(1, err)
	}
}

type RewinderFunc1[T any, E error] func(*Checkpoint[E], ...Option[E]) T

// Check1 checks whether err is not nil.
// If err is nil, it returns v.
// Otherwise it rewinds to the fallback point s, then [Handle] returns err.
//
// Check1 should be called on the same stack to [Handle].
func Check1[T any, E error](v T, err E) RewinderFunc1[T, E] {
	return func(cp *Checkpoint[E], opts ...Option[E]) T {
		applyOpts(cp, opts...)
		cp.raise(1, err)
		return v
	}
}

type RewinderFunc2[T1, T2 any, E error] func(*Checkpoint[E], ...Option[E]) (T1, T2)

// Check2 is a variant of [Check1].
func Check2[T1, T2 any, E error](v1 T1, v2 T2, err E) RewinderFunc2[T1, T2, E] {
	return func(cp *Checkpoint[E], opts ...Option[E]) (T1, T2) {
		applyOpts(cp, opts...)
		cp.raise(1, err)
		return v1, v2
	}
}
