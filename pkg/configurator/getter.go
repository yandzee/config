package configurator

import (
	"github.com/yandzee/config/pkg/source"
	"github.com/yandzee/config/pkg/str"
	"github.com/yandzee/config/pkg/str/transform"
)

type Getter struct {
	Source       source.StringSource
	ValueResults *[]*ValueResult[any]
}

func (g *Getter) Int(trs ...transform.Transformer) int {
	return NewGetterUnpacker[int]().
		Transformers(trs...).
		Transformers(transform.Parse(str.Parser.Int)).
		Unwrap(g)
}

func (g *Getter) IntOr(def int, trs ...transform.Transformer) int {
	return NewGetterUnpacker[int]().
		Transformers(trs...).
		Transformers(
			transform.Default(def),
			transform.Parse(str.Parser.Int),
		).
		Unwrap(g)
}

func (g *Getter) IntOrFn(fn transform.Defaulter[int], trs ...transform.Transformer) int {
	return NewGetterUnpacker[int]().
		Transformers(transform.DefaultFn(fn)).
		Transformers(trs...).
		Transformers(transform.Parse(str.Parser.Int)).
		Unwrap(g)
}

func (g *Getter) Bool(trs ...transform.Transformer) bool {
	return NewGetterUnpacker[bool]().
		Transformers(trs...).
		Transformers(transform.Parse(str.Parser.Bool)).
		Unwrap(g)
}

func (g *Getter) BoolOr(def bool, trs ...transform.Transformer) bool {
	return NewGetterUnpacker[bool]().
		Transformers(transform.Default(def)).
		Transformers(trs...).
		Transformers(transform.Parse(str.Parser.Bool)).
		Unwrap(g)
}

func (g *Getter) BoolOrFn(fn transform.Defaulter[bool], trs ...transform.Transformer) bool {
	return NewGetterUnpacker[bool]().
		Transformers(transform.DefaultFn(fn)).
		Transformers(trs...).
		Transformers(transform.Parse(str.Parser.Bool)).
		Unwrap(g)
}

func (g *Getter) Any(trs ...transform.Transformer) any {
	return NewGetterUnpacker[any]().
		Transformers(trs...).
		Unwrap(g)
}

func (g *Getter) AnyOr(def any, trs ...transform.Transformer) any {
	return NewGetterUnpacker[any]().
		Transformers(transform.Default(def)).
		Transformers(trs...).
		Unwrap(g)
}

func (g *Getter) AnyOrFn(fn transform.Defaulter[any], trs ...transform.Transformer) any {
	return NewGetterUnpacker[any]().
		Transformers(transform.DefaultFn(fn)).
		Transformers(trs...).
		Unwrap(g)
}
