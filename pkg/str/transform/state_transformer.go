package transform

type StateTransformer struct {
	Fn StateFn
}

func (st *StateTransformer) Chain(t Transformer) Transformer {
	return &StateTransformer{
		Fn: func(state *State) error {
			if err := st.Transform(state); err != nil {
				return err
			}

			return t.Transform(state)
		},
	}
}

func (st *StateTransformer) Transform(state *State) error {
	return st.Fn(state)
}
