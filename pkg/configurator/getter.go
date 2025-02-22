package configurator

import (
	"log/slog"

	"github.com/yandzee/config/pkg/common"
	"github.com/yandzee/config/pkg/source"
	"github.com/yandzee/config/pkg/str"
	"github.com/yandzee/config/pkg/str/transform"
)

// "chelnok-backend/pkg/config/options"
// "chelnok-backend/pkg/config/parse"

type Getter struct {
	// Configurator *Configurator
	Source        source.StringSource
	IsValueLogged bool
	LogRecords    *[]slog.Record
}

func (g *Getter) Bool(trs ...transform.Transformer) bool {
	return NewSourceParser[bool](g.Source).
		Transformers(trs...).
		Transformers(transform.Parse(str.Parser.Bool)).
		Unwrap(g)
}

func (g *Getter) BoolOr(def bool, trs ...transform.Transformer) bool {
	return NewSourceParser[bool](g.Source).
		Default(def).
		Transformers(trs...).
		Transformers(transform.Parse(str.Parser.Bool)).
		Unwrap(g)
}

func (g *Getter) BoolOrFn(fn common.DefaultFn[bool], trs ...transform.Transformer) bool {
	return NewSourceParser[bool](g.Source).
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(transform.Parse(str.Parser.Bool)).
		Unwrap(g)
}
