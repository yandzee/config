package str

import (
	"encoding/base64"
	"encoding/hex"
	"strings"

	"github.com/yandzee/config/pkg/str/parse"
	"github.com/yandzee/config/pkg/transform"
)

var (
	Parser = &parse.StringParser{}
)

func Parse[T any](fn parse.Fn[T]) transform.Transformer {
	return FromBytes().Chain(transform.Map(fn))
}

func Unbase64() transform.Transformer {
	return FromBytes().Chain(
		transform.Map(func(s string) (string, error) {
			decoded, err := base64.StdEncoding.DecodeString(s)
			return string(decoded), err
		}),
	)
}

func Unhex() transform.Transformer {
	return FromBytes().Chain(
		transform.Map(func(hexstr string) (string, error) {
			decoded, err := hex.DecodeString(strings.TrimPrefix(hexstr, "0x"))
			return string(decoded), err
		}),
	)
}

func Split(seps ...string) transform.Transformer {
	return FromBytes().Chain(
		transform.Map(func(str string) ([]string, error) {
			return Parser.Strings(str, seps...)
		}),
	)
}

func FromBytes() transform.Transformer {
	return transform.Map(func(strOrBytes any) (string, error) {
		switch casted := strOrBytes.(type) {
		case string:
			return casted, nil
		case []byte:
			return string(casted), nil
		}

		return "", transform.ErrConversion
	})
}
