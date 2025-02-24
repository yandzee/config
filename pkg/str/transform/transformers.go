package transform

import (
	"fmt"

	"github.com/yandzee/config/pkg/common"
	"github.com/yandzee/config/pkg/str/parse"
)

func Parse[T any](fn parse.Fn[T]) Transformer {
	return Map(fn)
}

func Map[F, T any](fn MapFromToFn[F, T]) Transformer {
	return StateTransform(func(s *State, opts common.KVOptions) error {
		val, ok := s.Value.(F)
		fmt.Printf("StateTransform: cast to From type: %v (%T) %v\n", val, val, ok)

		if !ok {
			return ErrCast
		}

		newVal, err := fn(val, opts)
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
