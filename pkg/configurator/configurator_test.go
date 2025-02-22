package configurator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

var (
	ErrTest1 = errors.New("ErrTest1")
	ErrTest2 = errors.New("ErrTest2")
)

type ConfiguratorTest[T any] struct {
	Action          func(*Configurator)
	ExpectedResults []ValueResult[any]
}

func TestConfigurator(t *testing.T) {
	runConfiguratorTests(t, []ConfiguratorTest[string]{
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4201", nil, true).Int()
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4201", nil, true),
					Value:  4201,
					Error:  nil,
					Flags:  DescFlagRequired | DescFlagPresented,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4202", nil, false).Bool()
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4202", nil, false),
					Value:  nil,
					Error:  nil,
					Flags:  DescFlagRequired | DescFlagNotPresented,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4203", ErrTest1, true).Int()
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4203", ErrTest1, true),
					Value:  nil,
					Error:  ErrTest1,
					Flags:  DescFlagRequired | DescFlagLookupError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4204", ErrTest2, false).Int()
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4204", ErrTest2, false),
					Value:  nil,
					Error:  nil,
					Flags:  DescFlagRequired | DescFlagNotPresented,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4205", nil, true).Bool()
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4205", nil, true),
					Value:  nil,
					Error:  strconv.ErrSyntax,
					Flags:  DescFlagRequired | DescFlagParseError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("T", nil, true).Bool()
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("T", nil, true),
					Value:  true,
					Error:  nil,
					Flags:  DescFlagRequired | DescFlagPresented,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("T", nil, true).BoolOr(true)
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("T", nil, true),
					Value:  true,
					Error:  nil,
					Flags:  DescFlagPresented,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("T", nil, false).BoolOr(true)
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("T", nil, false),
					Value:  true,
					Error:  nil,
					Flags:  DescFlagNotPresented | DescFlagDefaulted,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4206", nil, true).BoolOrFn(func() (bool, error) {
					return false, ErrTest1
				})
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4206", nil, true),
					Value:  nil,
					Error:  ErrTest1,
					Flags:  DescFlagPresented,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4207", nil, false).BoolOrFn(func() (bool, error) {
					return false, ErrTest1
				})
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4207", nil, false),
					Value:  nil,
					Error:  ErrTest1,
					Flags:  DescFlagPresented | DescFlagCustomError,
				},
			},
		},
	})
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

			logRecords := cfg.LogRecords()
			if len(logRecords) != ct.ExpectedLogRecords {
				t.Fatalf(
					"Num of log records is invalid, expected: %d, got: %d;\n%v\n",
					ct.ExpectedLogRecords,
					len(logRecords),
					logRecords,
				)
			}
		})
	}
}

func valptr[T any](val T) *T {
	return &val
}

func typename[T any]() string {
	return reflect.ValueOf(*new(T)).Kind().String()
}
