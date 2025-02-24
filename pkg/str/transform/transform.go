package transform

import (
	"errors"

	"github.com/yandzee/config/pkg/common"
)

var (
	ErrTransform = errors.New("Failed to transform data")
	ErrChain     = errors.New("Failed to chain transformers")

	ErrCast = errors.New("Failed to cast value type")
)

type State struct {
	Value any
}

type Transformer interface {
	Chain(Transformer) (Transformer, error)
	Transform(*State, common.KVOptions) error
}

type StateFn = func(*State, common.KVOptions) error
type MapFromToFn[F, T any] = func(F, common.KVOptions) (T, error)

// type MapFn = MapFromToFn[string, string]
// type MapSliceFn = MapFromToFn[[]string, []string]
// type MapIntoFn[T any] = MapFromToFn[string, T]
// type SplitFn = MapFromToFn[string, []string]
// type JoinFn = MapFromToFn[[]string, string]
