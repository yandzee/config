package configurator

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"testing"
)

var (
	ErrTest1 = errors.New("ErrTest1")
	ErrTest2 = errors.New("ErrTest2")
)

type ConfiguratorTest[T any] struct {
	Action          func(*Configurator)
	ExpectedResults []ExpectedResult
}

type ExpectedResult struct {
	Value    any
	Error    error
	Flags    DescriptorFlag
	LogLevel slog.Level
}

func TestConfigurator(t *testing.T) {
	runConfiguratorTests(t, []ConfiguratorTest[string]{
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4201", nil, true).Int()
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    4201,
					Error:    nil,
					Flags:    DescFlagRequired | DescFlagPresented,
					LogLevel: slog.LevelInfo,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4202", nil, false).Int()
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    nil,
					Flags:    DescFlagRequired,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4203", ErrTest1, true).Int()
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest1,
					Flags:    DescFlagRequired | DescFlagLookupError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4204", ErrTest2, false).Int()
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest2,
					Flags:    DescFlagRequired | DescFlagLookupError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("s4205", nil, true).Int()
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    strconv.ErrSyntax,
					Flags:    DescFlagRequired | DescFlagPresented | DescFlagParseError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4206", nil, true).IntOr(4207)
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    4206,
					Error:    nil,
					Flags:    DescFlagPresented,
					LogLevel: slog.LevelInfo,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4207", nil, false).IntOr(4208)
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    4208,
					Error:    nil,
					Flags:    DescFlagDefaulted,
					LogLevel: slog.LevelWarn,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4208", ErrTest1, true).IntOr(4209)
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest1,
					Flags:    DescFlagLookupError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4209", ErrTest2, false).IntOr(4210)
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest2,
					Flags:    DescFlagLookupError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4210", nil, true).IntOrFn(func() (int, error) {
					return 4211, nil
				})
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    4210,
					Error:    nil,
					Flags:    DescFlagPresented,
					LogLevel: slog.LevelInfo,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4211", nil, false).IntOrFn(func() (int, error) {
					return 4212, nil
				})
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    4212,
					Error:    nil,
					Flags:    DescFlagDefaulted,
					LogLevel: slog.LevelWarn,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4212", ErrTest1, true).IntOrFn(func() (int, error) {
					return 4213, nil
				})
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest1,
					Flags:    DescFlagLookupError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4213", ErrTest2, false).IntOrFn(func() (int, error) {
					return 4214, nil
				})
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest2,
					Flags:    DescFlagLookupError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *Configurator) {
				cfg.Str("4214", nil, false).IntOrFn(func() (int, error) {
					return 4215, ErrTest1
				})
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    4215,
					Error:    ErrTest1,
					Flags:    DescFlagDefaulted | DescFlagCustomError,
					LogLevel: slog.LevelError,
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

func checkValueResults[T any](
	t *testing.T,
	idx int,
	exp *ExpectedResult,
	got *ValueResult[any],
) {
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

	logEntry := got.LogRecord(false)
	hasValue := false

	logEntry.Attrs(func(a slog.Attr) bool {
		hasValue = hasValue || a.Key == LogAttrValue
		return true
	})

	if hasValue {
		t.Fatalf(
			"Value result %d: log record contains `value` even when not required",
			idx,
		)
	}

	if exp.LogLevel != 0 && exp.LogLevel != logEntry.Level {
		t.Fatalf(
			"Value result %d: log record levels are not equal, exp: %v, got: %v\n",
			idx,
			exp.LogLevel,
			logEntry.Level,
		)
	}
}
