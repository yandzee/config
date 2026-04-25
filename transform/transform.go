package transform

import (
	"errors"
	"fmt"
)

var (
	ErrConversion = errors.New("Failed to convert value type")
)

type Transformer interface {
	Chain(Transformer) Transformer
	Transform(any) (any, error)
}

type AnyToAnyFn func(any) (any, error)

func Map[F, T any](fn func(F) (T, error)) Transformer {
	return AnyToAny(func(val any) (any, error) {
		coerced, ok := val.(F)
		if !ok {
			return nil, errors.Join(
				ErrConversion,
				fmt.Errorf(
					"Failed to cast state value `%v` of type %T to type %T",
					val,
					val,
					coerced,
				),
			)
		}

		return fn(coerced)
	})
}

func AnyToAny(fn AnyToAnyFn) Transformer {
	return &FnTransformer{
		Fn: fn,
	}
}

func Run(val any, trs []Transformer) (any, error) {
	var err error

	for _, tr := range trs {
		val, err = tr.Transform(val)

		if err != nil {
			return nil, err
		}
	}

	return val, nil
}
