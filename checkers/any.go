package checkers

import "github.com/yandzee/config/check"

func AdaptChecker[T any](chckr check.Checker[any]) check.Checker[T] {
	return &AnyChecker[T]{
		Underlying: chckr,
	}
}

type AnyChecker[T any] struct {
	Underlying check.Checker[any]
}

func (ac *AnyChecker[T]) Check(v T) bool {
	return ac.Underlying.Check(any(v))
}
