package parse

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/yandzee/config/pkg/common"
)

type StringParser struct{}

func (sp *StringParser) Int(v string, opts common.KVOptions) (int, error) {
	parsed, err := strconv.ParseInt(v, 10, 0)
	return int(parsed), err
}

func (sp *StringParser) Int8(v string, opts common.KVOptions) (int8, error) {
	parsed, err := strconv.ParseInt(v, 10, 8)
	return int8(parsed), err
}

func (sp *StringParser) Int16(v string, opts common.KVOptions) (int16, error) {
	parsed, err := strconv.ParseInt(v, 10, 16)
	return int16(parsed), err
}

func (sp *StringParser) Int32(v string, opts common.KVOptions) (int32, error) {
	parsed, err := strconv.ParseInt(v, 10, 32)
	return int32(parsed), err
}

func (sp *StringParser) Int64(v string, opts common.KVOptions) (int64, error) {
	return strconv.ParseInt(v, 10, 64)
}

func (sp *StringParser) Uint(v string, opts common.KVOptions) (uint, error) {
	parsed, err := strconv.ParseUint(v, 10, 0)
	return uint(parsed), err
}

func (sp *StringParser) Uint8(v string, opts common.KVOptions) (uint8, error) {
	parsed, err := strconv.ParseUint(v, 10, 8)
	return uint8(parsed), err
}

func (sp *StringParser) Uint16(v string, opts common.KVOptions) (uint16, error) {
	parsed, err := strconv.ParseUint(v, 10, 16)
	return uint16(parsed), err
}

func (sp *StringParser) Uint32(v string, opts common.KVOptions) (uint32, error) {
	parsed, err := strconv.ParseUint(v, 10, 32)
	return uint32(parsed), err
}

func (sp *StringParser) Uint64(v string, opts common.KVOptions) (uint64, error) {
	return strconv.ParseUint(v, 10, 64)
}

func (sp *StringParser) Float32(v string, opts common.KVOptions) (float32, error) {
	parsed, err := strconv.ParseFloat(v, 32)
	return float32(parsed), err
}

func (sp *StringParser) Float64(v string, opts common.KVOptions) (float64, error) {
	return strconv.ParseFloat(v, 64)
}

func (sp *StringParser) Bool(v string, opts common.KVOptions) (bool, error) {
	return strconv.ParseBool(v)
}

func (sp *StringParser) Bytes(v string, opts common.KVOptions) ([]byte, error) {
	return []byte(v), nil
}

// func (sp *StringParser) String(v string, opts common.KVOptions) (string, error) {
// 	return v, nil
// }

func (sp *StringParser) Strings(v string, opts common.KVOptions) ([]string, error) {
	res := []string{v}
	separators := opts.GetStringSliceOr("seps", []string{","})
	trim := opts.GetBoolOr("trim", false)

	for _, sep := range separators {
		for i, str := range res {
			parts := strings.Split(str, sep)

			// Case when not splits occurred
			if len(parts) == 1 {
				if !trim {
					res[i] = parts[0]
				} else {
					res[i] = strings.TrimSpace(parts[0])
				}

				continue
			}

			// Move elems beginning from `i` to the right, make place for `parts`
			n := len(res) + len(parts) - 1
			resReplaced := make([]string, n)
			copy(resReplaced, res)

			for range len(parts) {
				resReplaced[n-1] = resReplaced[n-2]
			}

			// Actual insertion of `parts`
			for j := range len(parts) {
				if trim {
					resReplaced[i+j] = strings.TrimSpace(parts[j])
				} else {
					resReplaced[i+j] = parts[j]
				}
			}

			res = resReplaced
		}
	}

	return res, nil
}

func (sp *StringParser) Duration(v string, opts common.KVOptions) (time.Duration, error) {
	return time.ParseDuration(v)
}

func (sp *StringParser) ECPrivateKey(v string, opts common.KVOptions) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(v))
	if block == nil {
		return nil, fmt.Errorf("PEM block is not found")
	}

	pk, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.Join(
			fmt.Errorf("Failed to x509.ParseECPrivateKey"),
			err,
		)
	}

	return pk, nil
}

func (sp *StringParser) SlogLevel(v string, opts common.KVOptions) (slog.Level, error) {
	lvl := slog.LevelDebug
	err := lvl.UnmarshalText([]byte(v))

	return lvl, err
}
