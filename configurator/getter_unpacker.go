package configurator

import (
	"errors"
	"fmt"

	"github.com/yandzee/config/source"
	"github.com/yandzee/config/transform"
)

type GetterUnpacker[T any] struct {
	transformers []transform.Transformer
}

type Defaulter[T any] func() (T, error)

func NewGetterUnpacker[T any]() *GetterUnpacker[T] {
	return &GetterUnpacker[T]{}
}

func (sp *GetterUnpacker[T]) Transformers(trs ...transform.Transformer) *GetterUnpacker[T] {
	sp.transformers = append(sp.transformers, trs...)
	return sp
}

func (sp *GetterUnpacker[T]) Default(def T) *GetterUnpacker[T] {
	return sp.DefaultFn(func() (T, error) {
		return def, nil
	})
}

func (sp *GetterUnpacker[T]) DefaultFn(fn Defaulter[T]) *GetterUnpacker[T] {
	tr := transform.StateTransform(func(state transform.State) error {
		unpackState, ok := state.(*UnpackState)
		if !ok {
			return fmt.Errorf("Failed to coerce state to UnpackState")
		}

		unpackState.IsDefaulterSet = true
		if unpackState.IsInitialized {
			return nil
		}

		value, err := fn()
		if err != nil {
			unpackState.DefaulterError = err
			return err
		}

		unpackState.IsDefaulted = true
		unpackState.Value = value

		return nil
	})

	return sp.Transformers(tr)
}

func (sp *GetterUnpacker[T]) Unwrap(g *Getter) T {
	return sp.Parse(g).Value
}

func (sp *GetterUnpacker[T]) Parse(g *Getter) *ValueResult[T] {
	result := sp.ReadSource(g.Source)

	if g.ValueResults != nil {
		*g.ValueResults = append(*g.ValueResults, result.Any())
	}

	return result
}

func (sp *GetterUnpacker[T]) ReadSource(src source.StringSource) *ValueResult[T] {
	str, presented, err := src.Lookup()

	result := &ValueResult[T]{
		Source: src,
		Error:  err,
	}

	if err != nil {
		result.Flags.Add(DescFlagLookupError)
		return result
	}

	state := &UnpackState{
		IsInitialized: presented,
	}

	if presented {
		result.Flags.Add(DescFlagPresented)
		state.Value = str
	}

	result.Error = transform.Run(state, sp.transformers)

	if state.DefaulterError != nil {
		result.Error = state.DefaulterError
		result.Flags.Add(DescFlagCustomError)
	} else if result.Error != nil {
		result.Flags.Add(DescFlagTransformError)
	}

	if state.IsDefaulted {
		result.Flags.Add(DescFlagDefaulted)
	}

	if !state.IsDefaulterSet {
		result.Flags.Add(DescFlagRequired)
	}

	if result.Error != nil {
		return result
	}

	if !state.IsInitialized && !state.IsDefaulted {
		return result
	}

	ok := false
	result.Value, ok = state.Value.(T)
	if ok {
		return result
	}

	result.Flags.Add(DescFlagTransformError)
	result.Error = errors.Join(
		transform.ErrConversion,
		fmt.Errorf(
			"Failed to cast resulting value %v (%T) to type %T",
			state.Value,
			state.Value,
			result.Value,
		),
	)

	return result
}
