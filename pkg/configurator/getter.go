package configurator

import (
	"github.com/yandzee/config/pkg/common"
	"github.com/yandzee/config/pkg/source"
	"github.com/yandzee/config/pkg/str"
)

// "chelnok-backend/pkg/config/options"
// "chelnok-backend/pkg/config/parse"

type Getter struct {
	Configurator *Configurator
	Source       source.StringSource
}

func (g *Getter) Bool(opts ...str.Transformer) bool {
	return NewUnwrapper(str.Parser.Bool).Unwrap(g)
	// return p(parse.Bool).Options(opts).Unwrap(g)
}

func (g *Getter) BoolOr(def bool, opts ...any) bool {
	return false
	// return p(parse.Bool).Options(opts).Default(def).Unwrap(g)
}

func (g *Getter) BoolOrFn(fn common.DefaultFn[bool], opts ...any) bool {
	return false
	// return p(parse.Bool).Options(opts).DefaultFn(fn).Unwrap(g)
}

// func p[T any](parseFn parse.Fn[T]) *Parser[T] {
// 	return &Parser[T]{
// 		parseFn: parseFn,
// 	}
// }
