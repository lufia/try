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
	err E
}

type regs struct {
	sp    uintptr
	bp    uintptr
	ctxt  uintptr
	pc    uintptr
	probe uintptr // checks whether stack had been moved or not.
}

type options[E error] struct {
	handlers []func(err E) E
}

// Option configures [Check], [Check1] and [Check2].
type Option[E error] func(*options[E])

// WithHandler returns an option for a [Checkpoint] with f.
//
// When [Check] or its variants are called, f is invoked with the error if it is non-nil.
// Additionally, if f returns a zero value, [Check] or its variants do not treat it as an error.
func WithHandler[E error](f func(err E) E) Option[E] {
	return func(o *options[E]) {
		o.handlers = append(o.handlers, f)
	}
}

// WithDescription returns an option for a [Checkpoint] with [fmt.Printf]-style arguments.
func WithDescription(format string, args ...any) Option[error] {
	prefix := fmt.Sprintf(format, args...)
	return func(o *options[error]) {
		o.handlers = append(o.handlers, func(err error) error {
			return fmt.Errorf("%s: %w", prefix, err)
		})
	}
}

// WithIgnore returns an option for a [Checkpoint] to ignore errors.
func WithIgnore[E error](errs ...E) Option[E] {
	return func(o *options[E]) {
		o.handlers = append(o.handlers, func(err E) E {
			for _, e := range errs {
				if errors.Is(err, e) {
					var zero E
					return zero
				}
			}
			return err
		})
	}
}

func applyOpts[E error](o *options[E], opts ...Option[E]) {
	for _, opt := range opts {
		opt(o)
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

func (cp *Checkpoint[E]) raise(skip int, err E, o *options[E]) {
	if isZero(err) {
		return
	}
	for _, f := range o.handlers {
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

// Rewind rewinds current execution point to cp.
//
// Deprecated: use goto statement instead.
func (cp *Checkpoint[E]) Rewind(err E) {
	var o options[E]
	cp.raise(1, err, &o)
}

type RewinderFunc[E error] func(*Checkpoint[E], ...Option[E])

// Check checks whether err is not nil.
// If err is nil, it does nothing.
// Otherwise it rewinds to the fallback point cp, then [Handle] or [HandleFor] returns err.
//
// Check should be called on the same stack to [Handle] or [HandleFor].
func Check[E error](err E) RewinderFunc[E] {
	return func(cp *Checkpoint[E], opts ...Option[E]) {
		var o options[E]
		applyOpts(&o, opts...)
		cp.raise(1, err, &o)
	}
}

type RewinderFunc1[T any, E error] func(*Checkpoint[E], ...Option[E]) T

// Check1 checks whether err is not nil.
// If err is nil, it returns v.
// Otherwise it rewinds to the fallback point cp, then [Handle] or [HandleFor] returns err.
//
// Check1 should be called on the same stack to [Handle] or [HandleFor].
func Check1[T any, E error](v T, err E) RewinderFunc1[T, E] {
	return func(cp *Checkpoint[E], opts ...Option[E]) T {
		var o options[E]
		applyOpts(&o, opts...)
		cp.raise(1, err, &o)
		return v
	}
}

type RewinderFunc2[T1, T2 any, E error] func(*Checkpoint[E], ...Option[E]) (T1, T2)

// Check2 is a variant of [Check1].
func Check2[T1, T2 any, E error](v1 T1, v2 T2, err E) RewinderFunc2[T1, T2, E] {
	return func(cp *Checkpoint[E], opts ...Option[E]) (T1, T2) {
		var o options[E]
		applyOpts(&o, opts...)
		cp.raise(1, err, &o)
		return v1, v2
	}
}
