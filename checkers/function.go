package checkers

import "github.com/yandzee/config/check"

type FnChecker[T any] struct {
	fn func(T) (bool, string)
}

func (fc *FnChecker[T]) Check(v T) (bool, string) {
	return fc.fn(v)
}

func Fn[T any](fn func(T) (bool, string)) check.Checker[T] {
	return &FnChecker[T]{
		fn: fn,
	}
}
