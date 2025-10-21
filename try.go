// Package try provides error-handling utilities.
package try

import "fmt"

// Checkpoint represents the fallback point.
type Checkpoint struct {
	sp    uintptr
	bp    uintptr
	ctxt  uintptr
	pc    uintptr
	probe uintptr // BP of Handle's parent

	err     error
	handler func(err error) error
}

// Option configures [Check], [Check1] and [Check2].
type Option func(*Checkpoint)

func applyOpts(cp *Checkpoint, opts ...Option) {
	for _, opt := range opts {
		opt(cp)
	}
}

func WithHandler(f func(err error) error) Option {
	return func(cp *Checkpoint) {
		cp.handler = f
	}
}

func WithDescription(format string, args ...any) Option {
	prefix := fmt.Sprintf(format, args...)
	return func(cp *Checkpoint) {
		cp.handler = func(err error) error {
			return fmt.Errorf("%s: %w", prefix, err)
		}
	}
}

func waserror(cp *Checkpoint) bool
func raise(cp *Checkpoint) bool
func getbp(skip int) uintptr

// Handle creates a fallback point.
func Handle() (*Checkpoint, error) {
	var cp Checkpoint
	if waserror(&cp) {
		return nil, cp.err
	}
	return &cp, nil
}

func (cp *Checkpoint) raise(skip int, err error) {
	if err == nil {
		return
	}
	if cp.handler != nil {
		err = cp.handler(err)
	}
	cp.err = err

	bp := getbp(skip + 1)
	d := bp - cp.probe
	cp.probe += d
	cp.sp += d
	cp.bp += d
	cp.ctxt += d
	raise(cp)
	panic("do not reach here")
}

// Rewind rewinds current execution point to cp.
func (cp *Checkpoint) Rewind(err error) {
	cp.raise(1, err)
}

type RewinderFunc func(*Checkpoint, ...Option)

// Check checks whether err is not nil.
// If err is nil, it does nothing.
// Otherwise it rewinds to the fallback point s, then [Handle] returns err.
//
// Check should be called on the same stack to [Handle].
func Check(err error) RewinderFunc {
	return func(cp *Checkpoint, opts ...Option) {
		applyOpts(cp, opts...)
		cp.raise(1, err)
	}
}

type RewinderFunc1[T any] func(*Checkpoint, ...Option) T

// Check1 checks whether err is not nil.
// If err is nil, it returns v.
// Otherwise it rewinds to the fallback point s, then [Handle] returns err.
//
// Check1 should be called on the same stack to [Handle].
func Check1[T any](v T, err error) RewinderFunc1[T] {
	return func(cp *Checkpoint, opts ...Option) T {
		applyOpts(cp, opts...)
		cp.raise(1, err)
		return v
	}
}

type RewinderFunc2[T1, T2 any] func(*Checkpoint, ...Option) (T1, T2)

// Check2 is a variant of [Check1].
func Check2[T1, T2 any](v1 T1, v2 T2, err error) RewinderFunc2[T1, T2] {
	return func(cp *Checkpoint, opts ...Option) (T1, T2) {
		applyOpts(cp, opts...)
		cp.raise(1, err)
		return v1, v2
	}
}
