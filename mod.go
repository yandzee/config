package config

import (
	"log/slog"
	"time"

	c "github.com/yandzee/config/configurator"
	"github.com/yandzee/config/str"
	"github.com/yandzee/config/transform"
	"github.com/yandzee/config/transformers"
)

var Default = &c.Configurator{}

func Set[T any](target *T, opts ...any) *c.Getter[T] {
	g := defaultGetter[T]()
	g.Target = target

	switch any(target).(type) {
	case *int:
		g.Post(transformers.Parse(str.DefaultParser.Int))
	case *int8:
		g.Post(transformers.Parse(str.DefaultParser.Int8))
	case *int16:
		g.Post(transformers.Parse(str.DefaultParser.Int16))
	case *int32:
		g.Post(transformers.Parse(str.DefaultParser.Int32))
	case *int64:
		g.Post(transformers.Parse(str.DefaultParser.Int64))
	case *uint:
		g.Post(transformers.Parse(str.DefaultParser.Uint))
	case *uint8:
		g.Post(transformers.Parse(str.DefaultParser.Uint8))
	case *uint16:
		g.Post(transformers.Parse(str.DefaultParser.Uint64))
	case *uint32:
		g.Post(transformers.Parse(str.DefaultParser.Uint32))
	case *uint64:
		g.Post(transformers.Parse(str.DefaultParser.Uint64))
	case *float32:
		g.Post(transformers.Parse(str.DefaultParser.Float32))
	case *float64:
		g.Post(transformers.Parse(str.DefaultParser.Float64))
	case *complex64:
		g.Post(transformers.Parse(str.DefaultParser.Complex64))
	case *complex128:
		g.Post(transformers.Parse(str.DefaultParser.Complex128))
	case *bool:
		g.Post(transformers.Parse(str.DefaultParser.Bool))
	case *[]string:
		seps := transformers.CoerceOptions[string](opts)
		g.Post(transformers.Split(seps...))
	case *time.Duration:
		g.Post(transformers.Parse(str.DefaultParser.Duration))
	case *slog.Level:
		g.Post(transformers.Parse(str.DefaultParser.SlogLevel))
	case *[]byte:
		g.Post(transformers.ToBytes)
	}

	return g
}

func Int() *c.Getter[int] {
	return Set[int](nil)
}

func Int8() *c.Getter[int8] {
	return Set[int8](nil)
}

func Int16() *c.Getter[int16] {
	return Set[int16](nil)
}

func Int32() *c.Getter[int32] {
	return Set[int32](nil)
}

func Int64() *c.Getter[int64] {
	return Set[int64](nil)
}

func Uint() *c.Getter[uint] {
	return Set[uint](nil)
}

func Uint8() *c.Getter[uint8] {
	return Set[uint8](nil)
}

func Uint16() *c.Getter[uint16] {
	return Set[uint16](nil)
}

func Uint32() *c.Getter[uint32] {
	return Set[uint32](nil)
}

func Uint64() *c.Getter[uint64] {
	return Set[uint64](nil)
}

func Float32() *c.Getter[float32] {
	return Set[float32](nil)
}

func Float64() *c.Getter[float64] {
	return Set[float64](nil)
}

func Complex64() *c.Getter[complex64] {
	return Set[complex64](nil)
}

func Complex128() *c.Getter[complex128] {
	return Set[complex128](nil)
}

func Bool() *c.Getter[bool] {
	return Set[bool](nil)
}

func Duration() *c.Getter[time.Duration] {
	return Set[time.Duration](nil)
}

func SlogLevel() *c.Getter[slog.Level] {
	return Set[slog.Level](nil)
}

func String() *c.Getter[string] {
	return Set[string](nil)
}

func Strings(seps ...any) *c.Getter[[]string] {
	return Set[[]string](nil, seps...)
}

func Bytes() *c.Getter[[]byte] {
	return Set[[]byte](nil)
}

func Clear() {
	Default.Clear()
}

func defaultGetter[T any]() *c.Getter[T] {
	return &c.Getter[T]{
		Configurator: Default,
		Transformers: []transform.Transformer{},
	}
}
