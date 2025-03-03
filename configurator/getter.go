package configurator

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/yandzee/config/source"
	"github.com/yandzee/config/transform"
)

type Getter[T any] struct {
	Target       *T
	Configurator *Configurator
	Transformers []transform.Transformer
	Defaulter    Defaulter[T]
}

type Defaulter[T any] func() (T, error)

func (g *Getter[T]) Pre(trs ...transform.Transformer) *Getter[T] {
	g.Transformers = append(g.Transformers, trs...)
	return g
}

func (g *Getter[T]) Default(def T) *Getter[T] {
	return g.DefaultFn(func() (T, error) {
		return def, nil
	})
}

func (g *Getter[T]) DefaultFn(fn Defaulter[T]) *Getter[T] {
	g.Defaulter = fn
	return g
}

func (g *Getter[T]) Env(envVar string) T {
	return g.From(&source.EnvVarSource{
		VarName: envVar,
	})
}

func (g *Getter[T]) EnvOr(envVar string, def T) T {
	return g.From(&source.EnvVarSource{
		VarName: envVar,
	}, func() (T, error) {
		return def, nil
	})
}

func (g *Getter[T]) EnvOrFn(envVar string, defFn Defaulter[T]) T {
	return g.From(&source.EnvVarSource{
		VarName: envVar,
	}, defFn)
}

func (g *Getter[T]) From(src source.StringSource, def ...Defaulter[T]) T {
	result := g.TryFrom(src, def...)

	if g.Configurator == nil {
		lvl, msg := result.LevelAndMessage()

		if lvl == slog.LevelError {
			panic(fmt.Sprintf("%s: %s\n", src.Name(), msg))
		}
	}

	return result.Value
}

func (g *Getter[T]) TryFrom(src source.StringSource, def ...Defaulter[T]) *ValueResult[T] {
	str, presented, err := src.Lookup()

	result := &ValueResult[T]{
		Source: src,
		Error:  err,
	}

	if err != nil {
		result.Flags.Add(DescFlagLookupError)
		g.saveResult(result)

		return result
	}

	state := &UnpackState{
		IsInitialized: presented,
	}

	defaulter := g.getDefaulter(def...)
	if defaulter == nil {
		result.Flags.Add(DescFlagRequired)
	}

	switch {
	case presented:
		result.Flags.Add(DescFlagPresented)

		state.Value = str
		result.Error = transform.Run(state, g.Transformers)

		if result.Error != nil {
			result.Flags.Add(DescFlagTransformError)
		}
	case defaulter != nil:
		result.Value, result.Error = defaulter()

		if result.Error != nil {
			result.Flags.Add(DescFlagCustomError)
		} else {
			result.Flags.Add(DescFlagDefaulted)
		}

		fallthrough
	default:
		g.saveResult(result)
		return result
	}

	g.saveResult(result)
	if result.Error != nil {
		return result
	}
	//
	// if !state.IsInitialized && !state.IsDefaulted {
	// 	return result
	// }

	ok := false
	result.Value, ok = state.Value.(T)
	if ok {
		return result
	}

	result.Flags.Add(DescFlagTransformError)
	result.Error = errors.Join(
		transform.ErrConversion,
		fmt.Errorf(
			"Failed to coerce resulting value %v (%T) to type %T",
			state.Value,
			state.Value,
			result.Value,
		),
	)

	return result
}

func (g *Getter[T]) saveResult(res *ValueResult[T]) {
	if g.Configurator != nil {
		g.Configurator.ValueResults = append(g.Configurator.ValueResults, res.Any())
	}

	if g.Target != nil {
		*g.Target = res.Value
	}
}

func (g *Getter[T]) getDefaulter(def ...Defaulter[T]) Defaulter[T] {
	if g.Defaulter != nil {
		return g.Defaulter
	}

	for _, defFn := range def {
		if defFn != nil {
			return defFn
		}
	}

	return nil
}
