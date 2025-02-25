package transform

import (
	"fmt"

	"github.com/yandzee/config/pkg/str/parse"
)

type Defaulter[T any] = func() (T, error)

func Parse[T any](fn parse.Fn[T]) Transformer {
	return Map(fn)
}

func Default[T any](val T) Transformer {
	return StateTransform(func(s *State) error {
		if !s.Initialized {
			s.Value = s
			s.Defaulted = true
		}

		return nil
	})
}

func DefaultFn[T any](fn Defaulter[T]) Transformer {
	return StateTransform(func(s *State) error {
		val, err := fn()
		if err != nil {
			return err
		}

		if !s.Initialized {
			s.Value = val
			s.Defaulted = true
		}

		return nil
	})
}

func Split(separators ...string) Transformer {
	return Map(func(s string) ([]string, error) {
		return []string{}, nil
	})
}

func Map[F, T any](fn MapFromToFn[F, T]) Transformer {
	return StateTransform(func(s *State) error {
		if !s.Initialized {
			return nil
		}

		val, ok := s.Value.(F)
		fmt.Printf("StateTransform: cast to From type: %v (%T) %v\n", val, val, ok)

		if !ok {
			return ErrCast
		}

		newVal, err := fn(val)
		fmt.Printf("StateTransform: transformation result %v (%T), err: %v\n", newVal, newVal, err)
		if err != nil {
			return err
		}

		s.Value = newVal
		return nil
	})
}

func StateTransform(fn StateFn) Transformer {
	return &StateTransformer{
		Fn: fn,
	}
}
