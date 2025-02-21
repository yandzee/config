package configurator

import "github.com/yandzee/config/pkg/str/parse"

type Unwrapper[T any] struct {
	parseFn parse.Fn[T]
}

func NewUnwrapper[T any](fn parse.Fn[T]) *Unwrapper[T] {
	return &Unwrapper[T]{
		parseFn: fn,
	}
}

func (u *Unwrapper[T]) Unwrap(g *Getter) T {
	return *new(T)
}
