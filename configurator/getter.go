package configurator

import (
	"log/slog"
	"time"

	"github.com/yandzee/config/source"
	"github.com/yandzee/config/str"
	"github.com/yandzee/config/transform"
)

type Getter struct {
	Source       source.StringSource
	ValueResults *[]*ValueResult[any]
}

func (g *Getter) Int(trs ...transform.Transformer) int {
	return NewGetterUnpacker[int]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int)).
		Unwrap(g)
}

func (g *Getter) IntOr(def int, trs ...transform.Transformer) int {
	return NewGetterUnpacker[int]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int)).
		Unwrap(g)
}

func (g *Getter) IntOrFn(fn Defaulter[int], trs ...transform.Transformer) int {
	return NewGetterUnpacker[int]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int)).
		Unwrap(g)
}

func (g *Getter) Int8(trs ...transform.Transformer) int8 {
	return NewGetterUnpacker[int8]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int8)).
		Unwrap(g)
}

func (g *Getter) Int8Or(def int8, trs ...transform.Transformer) int8 {
	return NewGetterUnpacker[int8]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int8)).
		Unwrap(g)
}

func (g *Getter) Int8OrFn(fn Defaulter[int8], trs ...transform.Transformer) int8 {
	return NewGetterUnpacker[int8]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int8)).
		Unwrap(g)
}

func (g *Getter) Int16(trs ...transform.Transformer) int16 {
	return NewGetterUnpacker[int16]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int16)).
		Unwrap(g)
}

func (g *Getter) Int16Or(def int16, trs ...transform.Transformer) int16 {
	return NewGetterUnpacker[int16]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int16)).
		Unwrap(g)
}

func (g *Getter) Int16OrFn(fn Defaulter[int16], trs ...transform.Transformer) int16 {
	return NewGetterUnpacker[int16]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int16)).
		Unwrap(g)
}

func (g *Getter) Int32(trs ...transform.Transformer) int32 {
	return NewGetterUnpacker[int32]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int32)).
		Unwrap(g)
}

func (g *Getter) Int32Or(def int32, trs ...transform.Transformer) int32 {
	return NewGetterUnpacker[int32]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int32)).
		Unwrap(g)
}

func (g *Getter) Int32OrFn(fn Defaulter[int32], trs ...transform.Transformer) int32 {
	return NewGetterUnpacker[int32]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int32)).
		Unwrap(g)
}

func (g *Getter) Int64(trs ...transform.Transformer) int64 {
	return NewGetterUnpacker[int64]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int64)).
		Unwrap(g)
}

func (g *Getter) Int64Or(def int64, trs ...transform.Transformer) int64 {
	return NewGetterUnpacker[int64]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int64)).
		Unwrap(g)
}

func (g *Getter) Int64OrFn(fn Defaulter[int64], trs ...transform.Transformer) int64 {
	return NewGetterUnpacker[int64]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Int64)).
		Unwrap(g)
}

func (g *Getter) Uint(trs ...transform.Transformer) uint {
	return NewGetterUnpacker[uint]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint)).
		Unwrap(g)
}

func (g *Getter) UintOr(def uint, trs ...transform.Transformer) uint {
	return NewGetterUnpacker[uint]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint)).
		Unwrap(g)
}

func (g *Getter) UintOrFn(fn Defaulter[uint], trs ...transform.Transformer) uint {
	return NewGetterUnpacker[uint]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint)).
		Unwrap(g)
}

func (g *Getter) Uint8(trs ...transform.Transformer) uint8 {
	return NewGetterUnpacker[uint8]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint8)).
		Unwrap(g)
}

func (g *Getter) Uint8Or(def uint8, trs ...transform.Transformer) uint8 {
	return NewGetterUnpacker[uint8]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint8)).
		Unwrap(g)
}

func (g *Getter) Uint8OrFn(fn Defaulter[uint8], trs ...transform.Transformer) uint8 {
	return NewGetterUnpacker[uint8]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint8)).
		Unwrap(g)
}

func (g *Getter) Uint16(trs ...transform.Transformer) uint16 {
	return NewGetterUnpacker[uint16]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint16)).
		Unwrap(g)
}

func (g *Getter) Uint16Or(def uint16, trs ...transform.Transformer) uint16 {
	return NewGetterUnpacker[uint16]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint16)).
		Unwrap(g)
}

func (g *Getter) Uint16OrFn(fn Defaulter[uint16], trs ...transform.Transformer) uint16 {
	return NewGetterUnpacker[uint16]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint16)).
		Unwrap(g)
}

func (g *Getter) Uint32(trs ...transform.Transformer) uint32 {
	return NewGetterUnpacker[uint32]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint32)).
		Unwrap(g)
}

func (g *Getter) Uint32Or(def uint32, trs ...transform.Transformer) uint32 {
	return NewGetterUnpacker[uint32]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint32)).
		Unwrap(g)
}

func (g *Getter) Uint32OrFn(fn Defaulter[uint32], trs ...transform.Transformer) uint32 {
	return NewGetterUnpacker[uint32]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint32)).
		Unwrap(g)
}

func (g *Getter) Uint64(trs ...transform.Transformer) uint64 {
	return NewGetterUnpacker[uint64]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint64)).
		Unwrap(g)
}

func (g *Getter) Uint64Or(def uint64, trs ...transform.Transformer) uint64 {
	return NewGetterUnpacker[uint64]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint64)).
		Unwrap(g)
}

func (g *Getter) Uint64OrFn(fn Defaulter[uint64], trs ...transform.Transformer) uint64 {
	return NewGetterUnpacker[uint64]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Uint64)).
		Unwrap(g)
}

func (g *Getter) Float32(trs ...transform.Transformer) float32 {
	return NewGetterUnpacker[float32]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Float32)).
		Unwrap(g)
}

func (g *Getter) Float32Or(def float32, trs ...transform.Transformer) float32 {
	return NewGetterUnpacker[float32]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Float32)).
		Unwrap(g)
}

func (g *Getter) Float32OrFn(fn Defaulter[float32], trs ...transform.Transformer) float32 {
	return NewGetterUnpacker[float32]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Float32)).
		Unwrap(g)
}

func (g *Getter) Float64(trs ...transform.Transformer) float64 {
	return NewGetterUnpacker[float64]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Float64)).
		Unwrap(g)
}

func (g *Getter) Float64Or(def float64, trs ...transform.Transformer) float64 {
	return NewGetterUnpacker[float64]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Float64)).
		Unwrap(g)
}

func (g *Getter) Float64OrFn(fn Defaulter[float64], trs ...transform.Transformer) float64 {
	return NewGetterUnpacker[float64]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Float64)).
		Unwrap(g)
}

func (g *Getter) Bool(trs ...transform.Transformer) bool {
	return NewGetterUnpacker[bool]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Bool)).
		Unwrap(g)
}

func (g *Getter) BoolOr(def bool, trs ...transform.Transformer) bool {
	return NewGetterUnpacker[bool]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Bool)).
		Unwrap(g)
}

func (g *Getter) BoolOrFn(fn Defaulter[bool], trs ...transform.Transformer) bool {
	return NewGetterUnpacker[bool]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Bool)).
		Unwrap(g)
}

func (g *Getter) Bytes(trs ...transform.Transformer) []byte {
	return NewGetterUnpacker[[]byte]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Bytes)).
		Unwrap(g)
}

func (g *Getter) BytesOr(def []byte, trs ...transform.Transformer) []byte {
	return NewGetterUnpacker[[]byte]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Bytes)).
		Unwrap(g)
}

func (g *Getter) BytesOrFn(fn Defaulter[[]byte], trs ...transform.Transformer) []byte {
	return NewGetterUnpacker[[]byte]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Bytes)).
		Unwrap(g)
}

func (g *Getter) Duration(trs ...transform.Transformer) time.Duration {
	return NewGetterUnpacker[time.Duration]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Duration)).
		Unwrap(g)
}

func (g *Getter) DurationOr(def time.Duration, trs ...transform.Transformer) time.Duration {
	return NewGetterUnpacker[time.Duration]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Duration)).
		Unwrap(g)
}

func (g *Getter) DurationOrFn(fn Defaulter[time.Duration], trs ...transform.Transformer) time.Duration {
	return NewGetterUnpacker[time.Duration]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.Duration)).
		Unwrap(g)
}

func (g *Getter) SlogLevel(trs ...transform.Transformer) slog.Level {
	return NewGetterUnpacker[slog.Level]().
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.SlogLevel)).
		Unwrap(g)
}

func (g *Getter) SlogLevelOr(def slog.Level, trs ...transform.Transformer) slog.Level {
	return NewGetterUnpacker[slog.Level]().
		Default(def).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.SlogLevel)).
		Unwrap(g)
}

func (g *Getter) SlogLevelOrFn(fn Defaulter[slog.Level], trs ...transform.Transformer) slog.Level {
	return NewGetterUnpacker[slog.Level]().
		DefaultFn(fn).
		Transformers(trs...).
		Transformers(str.Parse(str.Parser.SlogLevel)).
		Unwrap(g)
}

func (g *Getter) Strings(separators ...string) []string {
	return NewGetterUnpacker[[]string]().
		Transformers(str.Split(separators...)).
		Unwrap(g)
}

func (g *Getter) StringsOr(def []string, separators ...string) []string {
	return NewGetterUnpacker[[]string]().
		Default(def).
		Transformers(str.Split(separators...)).
		Unwrap(g)
}

func (g *Getter) StringsOrFn(fn Defaulter[[]string], separators ...string) []string {
	return NewGetterUnpacker[[]string]().
		DefaultFn(fn).
		Transformers(str.Split(separators...)).
		Unwrap(g)
}

func (g *Getter) Any(trs ...transform.Transformer) any {
	return NewGetterUnpacker[any]().
		Transformers(trs...).
		Unwrap(g)
}

func (g *Getter) AnyOr(def any, trs ...transform.Transformer) any {
	return NewGetterUnpacker[any]().
		Default(def).
		Transformers(trs...).
		Unwrap(g)
}

func (g *Getter) AnyOrFn(fn Defaulter[any], trs ...transform.Transformer) any {
	return NewGetterUnpacker[any]().
		DefaultFn(fn).
		Transformers(trs...).
		Unwrap(g)
}
