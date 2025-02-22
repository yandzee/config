package configurator

import (
	"log/slog"

	"github.com/yandzee/config/pkg/common"
	"github.com/yandzee/config/pkg/source"
	"github.com/yandzee/config/pkg/str"
	"github.com/yandzee/config/pkg/str/transform"
)

type Getter struct {
	Source        source.StringSource
	IsValueLogged bool
	LogRecords    *[]slog.Record
}

func (g *Getter) Int(trs ...transform.Transformer) int {
	return NewGetterUnpacker[int]().
		Transformers(trs...).
		Transformers(transform.Parse(str.Parser.Int)).
		Unwrap(g)
}

func (g *Getter) IntOr(def int, trs ...transform.Transformer) int {
	return NewGetterUnpacker[int]().
		Default(def).
		Transformers(trs...).
		Transformers(transform.Parse(str.Parser.Int)).
		Unwrap(g)
}

func (g *Getter) IntOrFn(fn common.DefaultFn[int], trs ...transform.Transformer) int {
	return NewGetterUnpacker[int]().
		DefaultFn(fn).
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
		Default(def).
		Transformers(trs...).
		Transformers(transform.Parse(str.Parser.Bool)).
		Unwrap(g)
}

func (g *Getter) BoolOrFn(fn common.DefaultFn[bool], trs ...transform.Transformer) bool {
	return NewGetterUnpacker[bool]().
		DefaultFn(fn).
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
		Default(def).
		Transformers(trs...).
		Unwrap(g)
}

func (g *Getter) AnyOrFn(fn common.DefaultFn[any], trs ...transform.Transformer) any {
	return NewGetterUnpacker[any]().
		DefaultFn(fn).
		Transformers(trs...).
		Unwrap(g)
}
