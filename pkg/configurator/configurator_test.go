package configurator

import (
	"errors"
	"fmt"
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
				cfg.Str("4202", nil, false).Int()
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4202", nil, false),
					Value:  0,
					Error:  nil,
					Flags:  DescFlagRequired,
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
					Value:  0,
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
					Value:  0,
					Error:  ErrTest2,
					Flags:  DescFlagRequired | DescFlagLookupError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("s4205", nil, true).Int()
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("s4205", nil, true),
					Value:  0,
					Error:  strconv.ErrSyntax,
					Flags:  DescFlagRequired | DescFlagPresented | DescFlagParseError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4206", nil, true).IntOr(4207)
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4206", nil, true),
					Value:  4206,
					Error:  nil,
					Flags:  DescFlagPresented,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4207", nil, false).IntOr(4208)
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4207", nil, true),
					Value:  4208,
					Error:  nil,
					Flags:  DescFlagDefaulted,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4208", ErrTest1, true).IntOr(4209)
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4208", ErrTest1, true),
					Value:  0,
					Error:  ErrTest1,
					Flags:  DescFlagLookupError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4209", ErrTest2, false).IntOr(4210)
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4209", ErrTest2, true),
					Value:  0,
					Error:  ErrTest2,
					Flags:  DescFlagLookupError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4210", nil, true).IntOrFn(func() (int, error) {
					return 4211, nil
				})
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4210", nil, true),
					Value:  4210,
					Error:  nil,
					Flags:  DescFlagPresented,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4211", nil, false).IntOrFn(func() (int, error) {
					return 4212, nil
				})
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4211", nil, false),
					Value:  4212,
					Error:  nil,
					Flags:  DescFlagDefaulted,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4212", ErrTest1, true).IntOrFn(func() (int, error) {
					return 4213, nil
				})
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4212", ErrTest1, true),
					Value:  0,
					Error:  ErrTest1,
					Flags:  DescFlagLookupError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4213", ErrTest2, false).IntOrFn(func() (int, error) {
					return 4214, nil
				})
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4213", ErrTest2, false),
					Value:  0,
					Error:  ErrTest2,
					Flags:  DescFlagLookupError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4214", nil, false).IntOrFn(func() (int, error) {
					return 4215, ErrTest1
				})
			},
			ExpectedResults: []ValueResult[any]{
				{
					Source: NewStr("4214", nil, false),
					Value:  4215,
					Error:  ErrTest1,
					Flags:  DescFlagDefaulted | DescFlagCustomError,
				},
			},
		},
	})
}

func runConfiguratorTests[T any](t *testing.T, tests []ConfiguratorTest[T]) {
	for idx, ct := range tests {
		t.Run(fmt.Sprintf("Test %d", idx), func(t *testing.T) {
			cfg := Configurator{}

			ct.Action(&cfg)

			if expLen := len(ct.ExpectedResults); expLen != len(cfg.ValueResults) {
				t.Fatalf(
					"Value results amount doesnt match, expected %d, got: %d\n%v\n",
					expLen,
					len(cfg.ValueResults),
					cfg.ValueResults,
				)
			}

			for i, expResult := range ct.ExpectedResults {
				gotResult := cfg.ValueResults[i]

				checkValueResults[T](t, i, &expResult, gotResult)
			}
		})
	}
}

func checkValueResults[T any](t *testing.T, idx int, exp, got *ValueResult[any]) {
	if exp == nil || got == nil {
		t.Fatalf("Value result %d: some value results are nil\n", idx)
	}

	if exp.Flags != got.Flags {
		t.Fatalf(
			"Value result %d: flags are not equal, exp: %s (%v), got: %s (%v)\n",
			idx,
			exp.Flags,
			exp.Flags.Pairs(),
			got.Flags,
			got.Flags.Pairs(),
		)

	}

	if !errors.Is(got.Error, exp.Error) {
		t.Fatalf(
			"Value result %d: errors are not equal, exp: %v, got: %v\n",
			idx,
			exp.Error,
			got.Error,
		)
	}

	if exp.Value != got.Value {
		t.Fatalf(
			"Value result %d: values are not equal, exp: %v (%T), got: %v (%T)\n",
			idx,
			exp.Value,
			exp.Value,
			got.Value,
			got.Value,
		)
	}
}
