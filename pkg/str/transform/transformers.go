package transform

import "github.com/yandzee/config/pkg/common"

func StateTransform(fn StateTransformFn) Transformer {
	return &StateTransformer{
		Fn: fn,
	}
}

func Split(fn SplitTransformFn) Transformer {
	return StateTransform(func(s *State, opts common.KVOptions) error {
		str, err := s.GetStringValue()
		if err != nil {
			return err
		}

		strSlice, err := fn(str, opts)
		if err != nil {
			return err
		}

		s.Value = strSlice
		return nil
	})
}
