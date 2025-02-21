package parse

import "github.com/yandzee/config/pkg/common"

type Fn[T any] func(string, common.KVOptions) (T, error)
