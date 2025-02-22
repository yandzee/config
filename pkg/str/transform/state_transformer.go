package transform

import "github.com/yandzee/config/pkg/common"

type StateTransformer struct {
	Fn StateTransformFn
}

func (st *StateTransformer) Chain(t Transformer) (Transformer, error) {
	return &StateTransformer{
		Fn: func(s *State, opts common.KVOptions) error {
			if err := st.Transform(s, opts); err != nil {
				return err
			}

			return t.Transform(s, opts)
		},
	}, nil
}

func (st *StateTransformer) Transform(s *State, opts common.KVOptions) error {
	return nil
}
