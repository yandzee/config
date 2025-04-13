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
	target       *T
	configurator *Configurator
	transformers []transform.Transformer
	checkers     []check.Checker[T]
	defaulter    Defaulter[T]
}

type Defaulter[T any] func() (T, error)

func (g *Getter[T]) Pre(trs ...transform.Transformer) *Getter[T] {
	g.transformers = append(trs, g.transformers...)
	return g
}

func (g *Getter[T]) Post(trs ...transform.Transformer) *Getter[T] {
	g.transformers = append(g.transformers, trs...)
	return g
}

func (g *Getter[T]) Default(def T) *Getter[T] {
	return g.DefaultFn(func() (T, error) {
		return def, nil
	})
}

func (g *Getter[T]) DefaultFn(fn Defaulter[T]) *Getter[T] {
	g.defaulter = fn
	return g
}

func (g *Getter[T]) SetConfigurator(cfg *Configurator) *Getter[T] {
	g.configurator = cfg
	return g
}

func (g *Getter[T]) SetTarget(t *T) *Getter[T] {
	g.target = t
	return g
}

func (g *Getter[T]) Env(envVar string, def ...T) T {
	fns := make([]Defaulter[T], len(def))

	for idx, defValue := range def {
		fns[idx] = func() (T, error) {
			return defValue, nil
		}
	}

	return g.From(&source.EnvVarSource{
		VarName: envVar,
	}, fns...)
}

func (g *Getter[T]) EnvOr(envVar string, defFn Defaulter[T]) T {
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

	if g.configurator == nil {
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

	defaulter := g.getDefaulter(def...)
	if defaulter == nil {
		res.Flags.Add(result.FlagRequired)
	}

	var transformed any

	switch {
	case presented:
		res.Flags.Add(result.FlagPresented)

		transformed, res.Error = transform.Run(str, g.transformers)

		if res.Error != nil {
			res.Flags.Add(result.FlagTransformError)
			break
		}
	case defaulter != nil:
		var val T
		val, res.Error = defaulter()

		if res.Error != nil {
			res.Flags.Add(result.FlagDefaulterError)
		} else {
			res.Flags.Add(result.FlagDefaulted)
			res.Value = val
		}
	}

	if res.Error != nil {
		g.saveResult(res)
		return res
	}

	if presented {
		ok := false
		res.Value, ok = transformed.(T)

		if !ok {
			res.Flags.Add(result.FlagTransformError)
			res.Error = errors.Join(
				transform.ErrConversion,
				fmt.Errorf(
					"Failed to coerce resulting value %v (%T) to type %T",
					transformed,
					transformed,
					res.Value,
				),
			)
		}
	}

	for _, checker := range g.checkers {
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
	g.checkers = append(g.checkers, chkrs...)
	return g
}

func (g *Getter[T]) Check(fn func(*result.Result[T]) (bool, string)) *Getter[T] {
	g.checkers = append(g.checkers, checkers.Fn(fn))
	return g
}

func (g *Getter[T]) saveResult(res *result.Result[T]) {
	if g.configurator != nil {
		g.configurator.Results = append(g.configurator.Results, res.Any())
	}

	if g.target != nil {
		*g.target = res.Value
	}
}

func (g *Getter[T]) getDefaulter(def ...Defaulter[T]) Defaulter[T] {
	if g.defaulter != nil {
		return g.defaulter
	}

	for _, defFn := range def {
		if defFn != nil {
			return defFn
		}
	}

	return nil
}
