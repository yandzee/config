package configurator

import (
	"github.com/yandzee/config/pkg/common"
	"github.com/yandzee/config/pkg/source"
	"github.com/yandzee/config/pkg/str/transform"
)

type SourceParser[T any] struct {
	def          *T
	defFn        common.DefaultFn[T]
	transformers []transform.Transformer
}

func NewSourceParser[T any]() *SourceParser[T] {
	return &SourceParser[T]{}
}

func (sp *SourceParser[T]) Transformers(trs ...transform.Transformer) *SourceParser[T] {
	sp.transformers = append(sp.transformers, trs...)
	return sp
}

func (sp *SourceParser[T]) Default(def T) *SourceParser[T] {
	sp.def = &def
	return sp
}

func (sp *SourceParser[T]) DefaultFn(fn common.DefaultFn[T]) *SourceParser[T] {
	sp.defFn = fn
	return sp
}

func (sp *SourceParser[T]) Unwrap(g *Getter) T {
	result := sp.Parse(g.Source)

	if g.LogRecords != nil {
		*g.LogRecords = append(*g.LogRecords, result.LogRecord(g.IsValueLogged))
	}

	return result.Value
}

func (sp *SourceParser[T]) Parse(src source.StringSource) *ValueResult[T] {
	str, presented, err := src.Lookup()

	result := &ValueResult[T]{
		Source: src,
		Error:  err,
	}

	if err != nil {
		result.Flags.Add(DescFlagLookupError)
		return result
	}

	if sp.defFn == nil && sp.def == nil {
		result.Flags.Add(DescFlagRequired)
	} else {
		result.Flags.Add(DescFlagDefaulted)
	}

	if presented {
		var parseFnOpts common.KVOptions

		result.Flags.Add(DescFlagPresented)

		str, parseFnOpts, result.Error = options.Apply(sp.opts, str)

		if result.Error != nil {
			result.Flags.Add(DescFlagParseError)
			return result
		}

		result.Value, result.Error = sp.parseFn(str, parseFnOpts)

		if result.Error != nil {
			result.Flags.Add(DescFlagParseError)
		}
	} else {
		if sp.defaultFn != nil {
			result.Value, result.Error = sp.defaultFn()
		} else if sp.defaultValue != nil {
			result.Value = *sp.defaultValue
		}

		if result.Error != nil {
			result.Flags.Add(DescFlagCustomError)
			result.Flags.Remove(DescFlagDefaulted)
		}
	}

	return result
}
