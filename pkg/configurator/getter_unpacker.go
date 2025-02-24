package configurator

import (
	"github.com/yandzee/config/pkg/common"
	"github.com/yandzee/config/pkg/source"
	"github.com/yandzee/config/pkg/str/transform"
)

type GetterUnpacker[T any] struct {
	def          *T
	defFn        common.DefaultFn[T]
	transformers []transform.Transformer
}

func NewGetterUnpacker[T any]() *GetterUnpacker[T] {
	return &GetterUnpacker[T]{}
}

func (sp *GetterUnpacker[T]) Transformers(trs ...transform.Transformer) *GetterUnpacker[T] {
	sp.transformers = append(sp.transformers, trs...)
	return sp
}

func (sp *GetterUnpacker[T]) Default(def T) *GetterUnpacker[T] {
	sp.def = &def
	return sp
}

func (sp *GetterUnpacker[T]) DefaultFn(fn common.DefaultFn[T]) *GetterUnpacker[T] {
	sp.defFn = fn
	return sp
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

	if sp.defFn == nil && sp.def == nil {
		result.Flags.Add(DescFlagRequired)
	}

	if err != nil {
		result.Flags.Add(DescFlagLookupError)
		return result
	}

	if presented {
		result.Flags.Add(DescFlagPresented)

		state := &transform.State{
			Value: str,
		}

		for _, tr := range sp.transformers {
			result.Error = tr.Transform(state, common.KVOptions{})

			if result.Error != nil {
				result.Flags.Add(DescFlagParseError)
				return result
			}
		}

		ok := false
		result.Value, ok = state.Value.(T)

		if !ok {
			result.Flags.Add(DescFlagParseError)
			result.Error = transform.ErrCast
		}

		return result
	}

	if sp.defFn != nil {
		result.Value, result.Error = sp.defFn()
		result.Flags.Add(DescFlagDefaulted)
	} else if sp.def != nil {
		result.Value = *sp.def
		result.Flags.Add(DescFlagDefaulted)
	}

	if result.Error != nil {
		result.Flags.Add(DescFlagCustomError)
	}

	return result
}
