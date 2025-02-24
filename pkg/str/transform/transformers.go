package transform

import (
	"fmt"

	"github.com/yandzee/config/pkg/common"
	"github.com/yandzee/config/pkg/str/parse"
)

func Map(fn MapFn) Transformer {
	return MapFromTo(func(s string, opts common.KVOptions) (string, error) {
		return fn(s, opts)
	})
}

func MapSlice(fn MapSliceFn) Transformer {
	return MapFromTo(func(s []string, opts common.KVOptions) ([]string, error) {
		return fn(s, opts)
	})
}

func MapInto[T any](fn MapIntoFn[T]) Transformer {
	return MapFromTo(func(s string, opts common.KVOptions) (T, error) {
		return fn(s, opts)
	})
}

func Split(fn SplitFn) Transformer {
	return MapFromTo(func(s string, opts common.KVOptions) ([]string, error) {
		return fn(s, opts)
	})
}

func Join(fn JoinFn) Transformer {
	return MapFromTo(func(s []string, opts common.KVOptions) (string, error) {
		return fn(s, opts)
	})
}

func Parse[T any](fn parse.Fn[T]) Transformer {
	return MapInto(fn)
}

func MapFromTo[F, T any](fn MapFromToFn[F, T]) Transformer {
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
