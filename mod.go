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

func Of[T any](target *T, cfgrs ...*c.Configurator) *c.Getter[T] {
	g := defaultGetter[T](cfgrs...)
	g.Target = target

	switch any(target).(type) {
	case *int:
		g.Pre(transformers.Parse(str.DefaultParser.Int))
	case *int8:
		g.Pre(transformers.Parse(str.DefaultParser.Int8))
	case *int16:
		g.Pre(transformers.Parse(str.DefaultParser.Int16))
	case *int32:
		g.Pre(transformers.Parse(str.DefaultParser.Int32))
	case *int64:
		g.Pre(transformers.Parse(str.DefaultParser.Int64))
	case *uint:
		g.Pre(transformers.Parse(str.DefaultParser.Uint))
	case *uint8:
		g.Pre(transformers.Parse(str.DefaultParser.Uint8))
	case *uint16:
		g.Pre(transformers.Parse(str.DefaultParser.Uint64))
	case *uint32:
		g.Pre(transformers.Parse(str.DefaultParser.Uint32))
	case *uint64:
		g.Pre(transformers.Parse(str.DefaultParser.Uint64))
	case *float32:
		g.Pre(transformers.Parse(str.DefaultParser.Float32))
	case *float64:
		g.Pre(transformers.Parse(str.DefaultParser.Float64))
	case *complex64:
		g.Pre(transformers.Parse(str.DefaultParser.Complex64))
	case *complex128:
		g.Pre(transformers.Parse(str.DefaultParser.Complex128))
	case *bool:
		g.Pre(transformers.Parse(str.DefaultParser.Bool))
	case *time.Duration:
		g.Pre(transformers.Parse(str.DefaultParser.Duration))
	case *slog.Level:
		g.Pre(transformers.Parse(str.DefaultParser.SlogLevel))
	}

	return g
}

func Int(cfgrs ...*c.Configurator) *c.Getter[int] {
	return Of[int](nil, cfgrs...)
}

func Int8(cfgrs ...*c.Configurator) *c.Getter[int8] {
	return Of[int8](nil, cfgrs...)
}

func Int16(cfgrs ...*c.Configurator) *c.Getter[int16] {
	return Of[int16](nil, cfgrs...)
}

func Int32(cfgrs ...*c.Configurator) *c.Getter[int32] {
	return Of[int32](nil, cfgrs...)
}

func Int64(cfgrs ...*c.Configurator) *c.Getter[int64] {
	return Of[int64](nil, cfgrs...)
}

func Uint(cfgrs ...*c.Configurator) *c.Getter[uint] {
	return Of[uint](nil, cfgrs...)
}

func Uint8(cfgrs ...*c.Configurator) *c.Getter[uint8] {
	return Of[uint8](nil, cfgrs...)
}

func Uint16(cfgrs ...*c.Configurator) *c.Getter[uint16] {
	return Of[uint16](nil, cfgrs...)
}

func Uint32(cfgrs ...*c.Configurator) *c.Getter[uint32] {
	return Of[uint32](nil, cfgrs...)
}

func Uint64(cfgrs ...*c.Configurator) *c.Getter[uint64] {
	return Of[uint64](nil, cfgrs...)
}

func Float32(cfgrs ...*c.Configurator) *c.Getter[float32] {
	return Of[float32](nil, cfgrs...)
}

func Float64(cfgrs ...*c.Configurator) *c.Getter[float64] {
	return Of[float64](nil, cfgrs...)
}

func Complex64(cfgrs ...*c.Configurator) *c.Getter[complex64] {
	return Of[complex64](nil, cfgrs...)
}

func Complex128(cfgrs ...*c.Configurator) *c.Getter[complex128] {
	return Of[complex128](nil, cfgrs...)
}

func Bool(cfgrs ...*c.Configurator) *c.Getter[bool] {
	return Of[bool](nil, cfgrs...)
}

func Duration(cfgrs ...*c.Configurator) *c.Getter[time.Duration] {
	return Of[time.Duration](nil, cfgrs...)
}

func SlogLevel(cfgrs ...*c.Configurator) *c.Getter[slog.Level] {
	return Of[slog.Level](nil, cfgrs...)
}

func Clear() {
	Default.Clear()
}

func defaultGetter[T any](cfgrs ...*c.Configurator) *c.Getter[T] {
	configurator := Default

	for _, cfgr := range cfgrs {
		if cfgr == nil {
			continue
		}

		configurator = cfgr
		break
	}

	return &c.Getter[T]{
		Configurator: configurator,
		Transformers: []transform.Transformer{},
	}
}
