package mo

import (
	"errors"
	"fmt"
)

// Result type implementation
type Result[T any] struct {
	val T
	err error
}

// Ok result constructor
func Ok[T any](v T) Result[T] {
	return Result[T]{val: v}
}

// Err result constructor, err must be not nil
func Err[T any](e error) Result[T] {
	if e == nil {
		panic(errors.New("error must be not nil"))
	}
	ret := Result[T]{err: e}
	return ret
}

// OptionFrom makes a new Result from `val, err` statement
func ResultFrom[T any](v T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(v)
}

// IsOk: true is Ok, false if Err
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// Err is error getter
func (r Result[T]) Err() error {
	return r.err
}

// On adds more info (error context) to the Err as a text prefix
func (r Result[T]) On(on ...any) Result[T] {
	if r.err != nil && on != nil {
		r.err = fmt.Errorf("%s: %w", fmt.Sprint(on...), r.err)
	}
	return r
}

// Try returns T if Ok or panics if Err (can be intercepted by `Catch`),
// variadic `on` adds extra error context as a text prefix
func (r Result[T]) Try(on ...any) T {
	if r.err != nil {
		if on == nil {
			panic(fmt.Errorf("try err: %w", r.err))
		} else {
			panic(fmt.Errorf("%s: try err: %w", fmt.Sprint(on...), r.err))
		}
	}
	return r.val
}

// TryOr extracts value or provides defaultTry
func (r Result[T]) TryOr(defaultTry T) T {
	if r.err == nil {
		return r.val
	}
	return defaultTry
}

// Unpack converts Result to `val, err`
func (r Result[T]) Unpack() (T, error) {
	return r.val, r.err
}

func (r Result[T]) String() string {
	return fmt.Sprintf("Result{value:%+v, err:%v}", r.val, r.err)
}
