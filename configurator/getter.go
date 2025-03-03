package configurator

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/yandzee/config/check"
	"github.com/yandzee/config/checkers"
	"github.com/yandzee/config/result"
	"github.com/yandzee/config/source"
	"github.com/yandzee/config/transform"
)

type Getter[T any] struct {
	Target       *T
	Configurator *Configurator
	Transformers []transform.Transformer
	Checkers     []check.Checker[T]
	Defaulter    Defaulter[T]
}

type Defaulter[T any] func() (T, error)

func (g *Getter[T]) Pre(trs ...transform.Transformer) *Getter[T] {
	g.Transformers = append(trs, g.Transformers...)
	return g
}

func (g *Getter[T]) Post(trs ...transform.Transformer) *Getter[T] {
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

func (g *Getter[T]) SetConfigurator(cfg *Configurator) *Getter[T] {
	g.Configurator = cfg
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

func (g *Getter[T]) FromOr(src source.StringSource, def T) T {
	return g.From(src, func() (T, error) {
		return def, nil
	})
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

func (g *Getter[T]) TryFrom(src source.StringSource, def ...Defaulter[T]) *result.Result[T] {
	str, presented, err := src.Lookup()

	res := &result.Result[T]{
		Source: src,
		Error:  err,
	}

	if err != nil {
		res.Flags.Add(result.FlagLookupError)
		g.saveResult(res)

		return res
	}

	state := &UnpackState{
		IsInitialized: presented,
	}

	defaulter := g.getDefaulter(def...)
	if defaulter == nil {
		res.Flags.Add(result.FlagRequired)
	}

	switch {
	case presented:
		res.Flags.Add(result.FlagPresented)

		state.Value = str
		res.Error = transform.Run(state, g.Transformers)

		if res.Error != nil {
			res.Flags.Add(result.FlagTransformError)
			break
		}
	case defaulter != nil:
		var val T
		val, res.Error = defaulter()

		if res.Error != nil {
			res.Flags.Add(result.FlagCustomError)
		} else {
			res.Flags.Add(result.FlagDefaulted)
			res.Value = val
		}
	}

	if res.Error != nil {
		g.saveResult(res)
		return res
	}

	if state.IsInitialized {
		ok := false
		res.Value, ok = state.Value.(T)

		if !ok {
			res.Flags.Add(result.FlagTransformError)
			res.Error = errors.Join(
				transform.ErrConversion,
				fmt.Errorf(
					"Failed to coerce resulting value %v (%T) to type %T",
					state.Value,
					state.Value,
					res.Value,
				),
			)
		}
	}

	for _, checker := range g.Checkers {
		ok, desc := checker.Check(res)

		if !ok {
			res.Flags.Add(result.FlagCheckFailed)
			res.Error = fmt.Errorf("%s", desc)

			break
		}
	}

	g.saveResult(res)
	return res
}

func (g *Getter[T]) Checks(chkrs ...check.Checker[T]) *Getter[T] {
	g.Checkers = append(g.Checkers, chkrs...)
	return g
}

func (g *Getter[T]) Check(fn func(*result.Result[T]) (bool, string)) *Getter[T] {
	g.Checkers = append(g.Checkers, checkers.Fn(fn))
	return g
}

func (g *Getter[T]) saveResult(res *result.Result[T]) {
	if g.Configurator != nil {
		g.Configurator.Results = append(g.Configurator.Results, res.Any())
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
