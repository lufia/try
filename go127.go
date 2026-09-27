//go:build go1.27

package try

// Check checks whether err is not nil.
// If err is nil, it does nothing.
// Otherwise it rewinds to the fallback point cp, then [Handle] or [HandleFor] returns err.
//
// Check should be called on the same stack to [Handle] or [HandleFor].
func (cp *Checkpoint[E]) Check(err E) {
	var o options[E]
	cp.raise(1, err, &o)
}

// Check1 checks whether err is not nil.
// If err is nil, it returns v.
// Otherwise it rewinds to the fallback point cp, then [Handle] or [HandleFor] returns err.
//
// Check1 should be called on the same stack to [Handle] or [HandleFor].
func (cp *Checkpoint[E]) Check1[T any](v T, err E) T {
	var o options[E]
	cp.raise(1, err, &o)
	return v
}

// Check2 is a variant of [Checkpoint.Check1].
func (cp *Checkpoint[E]) Check2[T1, T2 any](v1 T1, v2 T2, err E) (T1, T2) {
	var o options[E]
	cp.raise(1, err, &o)
	return v1, v2
}

// CheckOptions is a variant of [Checkpoint.Check].
// It returns a function that accepts options.
func (cp *Checkpoint[E]) CheckOptions(err E) func(opts ...Option[E]) {
	return func(opts ...Option[E]) {
		var o options[E]
		applyOpts(&o, opts...)
		cp.raise(1, err, &o)
	}
}

// Check1Options is a variant of [Checkpoint.Check1].
// It returns a function that accepts options.
func (cp *Checkpoint[E]) Check1Options[T any](v T, err E) func(opts ...Option[E]) T {
	return func(opts ...Option[E]) T {
		var o options[E]
		applyOpts(&o, opts...)
		cp.raise(1, err, &o)
		return v
	}
}

// Check2Options is a variant of [Checkpoint.Check2].
// It returns a function that accepts options.
func (cp *Checkpoint[E]) Check2Options[T1, T2 any](v1 T1, v2 T2, err E) func(opts ...Option[E]) (T1, T2) {
	return func(opts ...Option[E]) (T1, T2) {
		var o options[E]
		applyOpts(&o, opts...)
		cp.raise(1, err, &o)
		return v1, v2
	}
}
