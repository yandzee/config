package checkers

import (
	"github.com/yandzee/config/check"
	"github.com/yandzee/config/result"
)

type FnChecker[T any] struct {
	fn func(*result.Result[T]) (bool, string)
}

func (fc *FnChecker[T]) Check(r *result.Result[T]) (bool, string) {
	return fc.fn(r)
}

func Fn[T any](fn func(*result.Result[T]) (bool, string)) check.Checker[T] {
	return &FnChecker[T]{
		fn: fn,
	}
}
