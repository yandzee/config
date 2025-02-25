package configurator

import (
	"github.com/yandzee/config/pkg/source"
	"github.com/yandzee/config/pkg/str"
	"github.com/yandzee/config/pkg/transform"
)

type Getter struct {
	Source       source.StringSource
	ValueResults *[]*ValueResult[any]
}

func (g *Getter) Int(trs ...transform.Transformer) int {
	return NewGetterUnpacker[int]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int)).
		Unwrap(g)
}

func (g *Getter) IntOr(def int, trs ...transform.Transformer) int {
	return NewGetterUnpacker[int]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int)).
		Unwrap(g)
}

func (g *Getter) IntOrFn(fn Defaulter[int], trs ...transform.Transformer) int {
	return NewGetterUnpacker[int]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int)).
		Unwrap(g)
}

func (g *Getter) Bool(trs ...transform.Transformer) bool {
	return NewGetterUnpacker[bool]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Bool)).
		Unwrap(g)
}

func (g *Getter) BoolOr(def bool, trs ...transform.Transformer) bool {
	return NewGetterUnpacker[bool]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Bool)).
		Unwrap(g)
}

func (g *Getter) BoolOrFn(fn Defaulter[bool], trs ...transform.Transformer) bool {
	return NewGetterUnpacker[bool]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Bool)).
		Unwrap(g)
}

func (g *Getter) Any(trs ...transform.Transformer) any {
	return NewGetterUnpacker[any]().
		Transformers(trs...).
		Unwrap(g)
}

func (g *Getter) AnyOr(def any, trs ...transform.Transformer) any {
	return NewGetterUnpacker[any]().
		Default(def).
		Transformers(trs...).
		Unwrap(g)
}

func (g *Getter) AnyOrFn(fn Defaulter[any], trs ...transform.Transformer) any {
	return NewGetterUnpacker[any]().
		DefaultFn(fn).
		Transformers(trs...).
		Unwrap(g)
}
