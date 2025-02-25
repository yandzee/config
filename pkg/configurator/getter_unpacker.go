package configurator

import (
	"github.com/yandzee/config/pkg/source"
	"github.com/yandzee/config/pkg/str/transform"
)

type GetterUnpacker[T any] struct {
	// def          *T
	// defFn        common.DefaultFn[T]
	transformers []transform.Transformer
}

func NewGetterUnpacker[T any]() *GetterUnpacker[T] {
	return &GetterUnpacker[T]{}
}

func (sp *GetterUnpacker[T]) Transformers(trs ...transform.Transformer) *GetterUnpacker[T] {
	sp.transformers = append(sp.transformers, trs...)
	return sp
}

// func (sp *GetterUnpacker[T]) Default(def T) *GetterUnpacker[T] {
// 	sp.def = &def
// 	return sp
// }
//
// func (sp *GetterUnpacker[T]) DefaultFn(fn common.DefaultFn[T]) *GetterUnpacker[T] {
// 	sp.defFn = fn
// 	return sp
// }

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

	if err != nil {
		result.Flags.Add(DescFlagLookupError)
		return result
	}

	state := &transform.State{
		Initialized: presented,
		Defaulted:   false,
	}

	if presented {
		result.Flags.Add(DescFlagPresented)

		state.Value = str
	}

	for _, tr := range sp.transformers {
		result.Error = tr.Transform(state)

		if result.Error != nil {
			result.Flags.Add(DescFlagTransformError)
			break
		}
	}

	if state.Defaulted {
		result.Flags.Add(DescFlagDefaulted)
	} else {
		result.Flags.Add(DescFlagRequired)
	}

	if result.Error != nil {
		return result
	}

	if state.Initialized {
		ok := false
		result.Value, ok = state.Value.(T)

		if !ok {
			result.Flags.Add(DescFlagTransformError)
			result.Error = transform.ErrCast

			return result
		}
	}

	// return result

	// if sp.defFn != nil {
	// 	result.Value, result.Error = sp.defFn()
	// 	result.Flags.Add(DescFlagDefaulted)
	// } else if sp.def != nil {
	// 	result.Value = *sp.def
	// 	result.Flags.Add(DescFlagDefaulted)
	// }

	// if result.Error != nil {
	// 	result.Flags.Add(DescFlagCustomError)
	// }

	return result
}
