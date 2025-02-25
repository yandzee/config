package transform

import (
	"errors"
	"fmt"
)

var (
	ErrTransform = errors.New("Failed to transform data")
	ErrCast      = errors.New("Failed to cast value type")

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
type MapFromToFn[F, T any] = func(F) (T, error)

func Map[F, T any](fn MapFromToFn[F, T]) Transformer {
	return StateTransform(func(state State) error {
		stateValue, err := state.GetValue()

		switch {
		case errors.Is(err, ErrNoValue):
			return nil
		case err != nil:
			return err
		}

		value, ok := stateValue.(F)
		fmt.Printf(
			"StateTransform: cast %v (%T) to type %T: %v\n",
			stateValue,
			stateValue,
			value,
			value,
		)

		if !ok {
			return errors.Join(
				ErrCast,
				fmt.Errorf(
					"Failed to cast state value `%v` of type %T to type %T",
					stateValue,
					stateValue,
					value,
				),
			)
		}

		newValue, err := fn(value)
		fmt.Printf(
			"StateTransform: transformation result, value: %v (%T), err: %v\n",
			newValue,
			newValue,
			err,
		)

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
