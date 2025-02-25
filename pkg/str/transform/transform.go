package transform

import (
	"errors"
)

var (
	ErrTransform = errors.New("Failed to transform data")

	ErrCast = errors.New("Failed to cast value type")
)

type State struct {
	Value       any
	Initialized bool
	Defaulted   bool
}

type Transformer interface {
	Chain(Transformer) Transformer
	Transform(*State) error
}

type StateFn = func(*State) error
type MapFromToFn[F, T any] = func(F) (T, error)

// type MapFn = MapFromToFn[string, string]
// type MapSliceFn = MapFromToFn[[]string, []string]
// type MapIntoFn[T any] = MapFromToFn[string, T]
// type SplitFn = MapFromToFn[string, []string]
// type JoinFn = MapFromToFn[[]string, string]
