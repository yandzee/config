package str

import (
	"errors"

	"github.com/yandzee/config/pkg/common"
)

var (
	ErrTransform = errors.New("Failed to transform data")
	ErrChain     = errors.New("Failed to chain transformers")
)

type Transformable struct {
	String  string
	Slice   []string
	IsSlice bool
}

type Transformer interface {
	Chain(Transformer) (Transformer, error)
	Transform(*Transformable, common.KVOptions) (*Transformable, error)
}
