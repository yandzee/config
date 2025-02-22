package str

import (
	"github.com/yandzee/config/pkg/common"
	"github.com/yandzee/config/pkg/str/parse"
	"github.com/yandzee/config/pkg/str/transform"
)

var (
	Parser = &parse.StringParser{}
)

var Split = transform.Split(func(s string, opts common.KVOptions) ([]string, error) {
	return Parser.Strings(s, opts)
})
