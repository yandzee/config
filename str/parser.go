package str

import (
	"log/slog"
	"slices"
	"strconv"
	"strings"
	"time"
)

var (
	DefaultParser = &StringParser{}
)

type StringParser struct{}

type ParseFn[T any] func(string) (T, error)

func (sp *StringParser) Int(v string) (int, error) {
	parsed, err := strconv.ParseInt(v, 10, 0)
	return int(parsed), err
}

func (sp *StringParser) Int8(v string) (int8, error) {
	parsed, err := strconv.ParseInt(v, 10, 8)
	return int8(parsed), err
}

func (sp *StringParser) Int16(v string) (int16, error) {
	parsed, err := strconv.ParseInt(v, 10, 16)
	return int16(parsed), err
}

func (sp *StringParser) Int32(v string) (int32, error) {
	parsed, err := strconv.ParseInt(v, 10, 32)
	return int32(parsed), err
}

func (sp *StringParser) Int64(v string) (int64, error) {
	return strconv.ParseInt(v, 10, 64)
}

func (sp *StringParser) Uint(v string) (uint, error) {
	parsed, err := strconv.ParseUint(v, 10, 0)
	return uint(parsed), err
}

func (sp *StringParser) Uint8(v string) (uint8, error) {
	parsed, err := strconv.ParseUint(v, 10, 8)
	return uint8(parsed), err
}

func (sp *StringParser) Uint16(v string) (uint16, error) {
	parsed, err := strconv.ParseUint(v, 10, 16)
	return uint16(parsed), err
}

func (sp *StringParser) Uint32(v string) (uint32, error) {
	parsed, err := strconv.ParseUint(v, 10, 32)
	return uint32(parsed), err
}

func (sp *StringParser) Uint64(v string) (uint64, error) {
	return strconv.ParseUint(v, 10, 64)
}

func (sp *StringParser) Float32(v string) (float32, error) {
	parsed, err := strconv.ParseFloat(v, 32)
	return float32(parsed), err
}

func (sp *StringParser) Float64(v string) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

func (sp *StringParser) Complex64(v string) (complex64, error) {
	parsed, err := strconv.ParseComplex(v, 64)
	return complex64(parsed), err
}

func (sp *StringParser) Complex128(v string) (complex128, error) {
	return strconv.ParseComplex(v, 128)
}

func (sp *StringParser) Bool(v string) (bool, error) {
	return strconv.ParseBool(v)
}

func (sp *StringParser) Bytes(v string) ([]byte, error) {
	return []byte(v), nil
}

func (sp *StringParser) Strings(v string, seps ...string) ([]string, error) {
	res := []string{v}

	for _, sep := range seps {
		for i := 0; i < len(res); i += 1 {
			str := res[i]

			parts := strings.Split(str, sep)
			n := len(parts)

			if n == 1 {
				continue
			}

			res = slices.Grow(res, n-1)

			res[i] = parts[0]
			if n > 1 {
				res = slices.Insert(res, i+1, parts[1:]...)
				i += n - 1
			}
		}
	}

	return res, nil
}

func (sp *StringParser) Duration(v string) (time.Duration, error) {
	return time.ParseDuration(v)
}

func (sp *StringParser) SlogLevel(v string) (slog.Level, error) {
	lvl := slog.LevelDebug
	err := lvl.UnmarshalText([]byte(v))

	return lvl, err
}
