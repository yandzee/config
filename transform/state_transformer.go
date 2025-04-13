package transform

type FnTransformer struct {
	Fn ValueTransformerFn
}

func (ft *FnTransformer) Chain(rhs Transformer) Transformer {
	return &FnTransformer{
		Fn: func(val any) (any, error) {
			transformed, err := ft.Transform(val)
			if err != nil {
				return nil, err
			}

			return rhs.Transform(transformed)
		},
	}
}

func (st *FnTransformer) Transform(val any) (any, error) {
	return st.Fn(val)
}
