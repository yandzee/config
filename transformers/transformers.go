package transformers

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"

	"github.com/yandzee/config/str"
	"github.com/yandzee/config/transform"
)

var (
	ToString = transform.Map(func(smth any) (string, error) {
		switch casted := smth.(type) {
		case string:
			return casted, nil
		case []byte:
			return string(casted), nil
		}

		return "", errors.Join(
			transform.ErrConversion,
			fmt.Errorf("Cannot convert %v (%T) into string", smth, smth),
		)
	})

	ToBytes = transform.Map(func(smth any) ([]byte, error) {
		switch casted := smth.(type) {
		case string:
			return []byte(casted), nil
		case []byte:
			return casted, nil
		}

		return nil, errors.Join(
			transform.ErrConversion,
			fmt.Errorf("Cannot convert %v (%T) into []byte", smth, smth),
		)
	})

	Unbase64 = ToString.Chain(
		transform.Map(func(s string) (string, error) {
			decoded, err := base64.StdEncoding.DecodeString(s)
			return string(decoded), err
		}),
	)

	Unhex = ToString.Chain(
		transform.Map(func(hexstr string) ([]byte, error) {
			return hex.DecodeString(strings.TrimPrefix(hexstr, "0x"))
		}),
	)

	ECPrivateKey = ToBytes.Chain(
		transform.Map(func(b []byte) (*ecdsa.PrivateKey, error) {
			block, _ := pem.Decode(b)
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
		}),
	)
)

func Parse[T any](fn str.ParseFn[T]) transform.Transformer {
	return ToString.Chain(transform.Map(fn))
}

func Split(seps ...string) transform.Transformer {
	return ToString.Chain(
		transform.Map(func(s string) ([]string, error) {
			return str.DefaultParser.Strings(s, seps...)
		}),
	)
}
