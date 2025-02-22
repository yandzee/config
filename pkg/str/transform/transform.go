package transform

import (
	"errors"

	"github.com/yandzee/config/pkg/common"
)

var (
	ErrTransform = errors.New("Failed to transform data")
	ErrChain     = errors.New("Failed to chain transformers")

	ErrNotString = errors.New("Failed to convert value to string")
	ErrNotSlice  = errors.New("Failed to convert value to string slice")
)

type State struct {
	Value any
}

type Transformer interface {
	Chain(Transformer) (Transformer, error)
	Transform(*State, common.KVOptions) error
}

type SplitTransformFn = func(string, common.KVOptions) ([]string, error)
type JoinTransformFn = func([]string, common.KVOptions) (string, error)
type StateTransformFn = func(*State, common.KVOptions) error

func (s *State) GetStringValue() (string, error) {
	str, ok := s.Value.(string)
	if !ok {
		return "", ErrNotString
	}

	return str, nil
}
