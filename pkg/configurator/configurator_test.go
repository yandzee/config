package configurator

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yandzee/config/pkg/common"
	"github.com/yandzee/config/pkg/source"
)

type ConfiguratorTest[T any] struct {
	Input              source.StringSource
	Default            *T
	DefaultFn          common.DefaultFn[T]
	Expected           *T
	ExpectedLogRecords int
}

func TestConfigurator(t *testing.T) {
	runConfiguratorTests(t, []ConfiguratorTest[string]{
		{
			Input: &source.StrSource{
				Str:       "42",
				Presented: true,
			},
			Expected:           valptr("42"),
			ExpectedLogRecords: 1,
		},
		{
			Input: &source.StrSource{
				Str:       "42",
				Presented: false,
			},
			Expected:           nil,
			ExpectedLogRecords: 1,
		},
	})
}

func valptr[T any](val T) *T {
	return &val
}

func typename[T any]() string {
	return reflect.ValueOf(*new(T)).Kind().String()
}

func runConfiguratorTests[T any](t *testing.T, tests []ConfiguratorTest[T]) {
	for idx, ct := range tests {
		t.Run(fmt.Sprintf("Test %d", idx), func(t *testing.T) {
			cfg := Configurator{}

			getter := cfg.Source(ct.Input)
			var output any

			switch {
			case ct.Default != nil:
				output = getter.AnyOr(*ct.Default)
			case ct.DefaultFn != nil:
				fn := func() (any, error) {
					return ct.DefaultFn()
				}

				output = getter.AnyOrFn(fn)
			default:
				output = getter.Any()
			}

			if ct.Expected != nil && output != nil {
				casted, ok := output.(T)
				if !ok {
					t.Fatalf(
						"Failed to coerce value: %v (%T) to %s, (is nil: %v)\n",
						output,
						output,
						typename[T](),
						output == nil,
					)
				}

				if !reflect.DeepEqual(*ct.Expected, casted) {
					t.Fatalf(
						"Expected output: %v (%T), got: %v (%T)\n",
						*ct.Expected,
						*ct.Expected,
						casted,
						casted,
					)
				}
			} else if ct.Expected == nil && output != nil {
				t.Fatalf("Expected value is nil, but output is %v (%T)\n", output, output)
			} else if ct.Expected != nil && output == nil {
				t.Fatalf("Expected %v (%T), but output is nil\n", *ct.Expected, *ct.Expected)
			}

			if len(cfg.LogRecords) != ct.ExpectedLogRecords {
				t.Fatalf(
					"Num of log records is invalid, expected: %d, got: %d;\n%v\n",
					ct.ExpectedLogRecords,
					len(cfg.LogRecords),
					cfg.LogRecords,
				)
			}
		})
	}
}
