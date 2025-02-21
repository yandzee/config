package str

import "github.com/yandzee/config/pkg/str/parse"

var (
	Parser = &parse.StringParser{}
)

var Split = Transform("Split", func(s string, opts common.KVOptions) (*Transformable, error) {
	strs, err := Parser.Strings(s, opts)
	if err != nil {
		return nil, err
	}

	return &Transformable{
		IsSlice: true,
		Slice:   strs,
	}, nil
})
