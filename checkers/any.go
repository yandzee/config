package checkers

import (
	"github.com/yandzee/config/check"
	"github.com/yandzee/config/result"
)

func AdaptChecker[T any](chckr check.Checker[any]) check.Checker[T] {
	return &AnyChecker[T]{
		Underlying: chckr,
	}
}

type AnyChecker[T any] struct {
	Underlying check.Checker[any]
}

func (ac *AnyChecker[T]) Check(r *result.Result[T]) (bool, string) {
	return ac.Underlying.Check(r.Any())
}
