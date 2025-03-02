package transform

import (
	"errors"
	"fmt"
)

var (
	ErrTransform  = errors.New("Failed to transform data")
	ErrConversion = errors.New("Failed to convert value type")

	ErrNoValue = errors.New("Transformable state has no value")
)

type State interface {
	GetValue() (any, error)
	SetValue(any) error
}

type Transformer interface {
	Chain(Transformer) Transformer
	Transform(State) error
}

type StateFn = func(State) error
// type MapFromToFn[F, T any] func(F) (T, error)

func Map[F, T any](fn func(F) (T, error)) Transformer {
	return StateTransform(func(state State) error {
		stateValue, err := state.GetValue()

		switch {
		case errors.Is(err, ErrNoValue):
			return nil
		case err != nil:
			return err
		}

		value, ok := stateValue.(F)
		if !ok {
			return errors.Join(
				ErrConversion,
				fmt.Errorf(
					"Failed to cast state value `%v` of type %T to type %T",
					stateValue,
					stateValue,
					value,
				),
			)
		}

		newValue, err := fn(value)
		if err != nil {
			return err
		}

		return state.SetValue(newValue)
	})
}

func StateTransform(fn StateFn) Transformer {
	return &StateTransformer{
		Fn: fn,
	}
}

func Run(state State, trs []Transformer) error {
	for _, tr := range trs {
		if err := tr.Transform(state); err != nil {
			return err
		}
	}

	return nil
}
