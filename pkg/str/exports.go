package str

import (
	"github.com/yandzee/config/pkg/str/parse"
	"github.com/yandzee/config/pkg/transform"
)

var (
	Parser = &parse.StringParser{}
)

func Parse[T any](fn parse.Fn[T]) transform.Transformer {
	return transform.Map(fn)
}
