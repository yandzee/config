package tests

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"testing"

	"github.com/yandzee/config"
	"github.com/yandzee/config/configurator"
	"github.com/yandzee/config/result"
)

var (
	ErrTest1 = errors.New("ErrTest1")
	ErrTest2 = errors.New("ErrTest2")
)

type ConfiguratorTest[T any] struct {
	Action          func(*configurator.Configurator)
	ExpectedResults []ExpectedResult
}

type ExpectedResult struct {
	Value    any
	Error    error
	ErrorFn  func(error) bool
	Flags    result.ResultFlag
	LogLevel slog.Level
}

func TestConfigurator(t *testing.T) {
	runConfiguratorTests(t, []ConfiguratorTest[string]{
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).From(NewStr("4201", nil, true))
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    4201,
					Error:    nil,
					Flags:    result.FlagRequired | result.FlagPresented,
					LogLevel: slog.LevelInfo,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).From(NewStr("4202", nil, false))
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    nil,
					Flags:    result.FlagRequired,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).From(NewStr("4203", ErrTest1, true))
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest1,
					Flags:    result.FlagLookupError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).From(NewStr("4204", ErrTest2, false))
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest2,
					Flags:    result.FlagLookupError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).From(NewStr("s4205", nil, true))
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    strconv.ErrSyntax,
					Flags:    result.FlagRequired | result.FlagPresented | result.FlagTransformError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).FromOr(NewStr("4206", nil, true), 4207)
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    4206,
					Error:    nil,
					Flags:    result.FlagPresented,
					LogLevel: slog.LevelInfo,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).FromOr(NewStr("4207", nil, false), 4208)
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    4208,
					Error:    nil,
					Flags:    result.FlagDefaulted,
					LogLevel: slog.LevelWarn,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).FromOr(NewStr("4208", ErrTest1, true), 4209)
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest1,
					Flags:    result.FlagLookupError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).FromOr(NewStr("4209", ErrTest2, false), 4210)
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest2,
					Flags:    result.FlagLookupError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).From(NewStr("4210", nil, true), func() (int, error) {
					return 4211, nil
				})
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    4210,
					Error:    nil,
					Flags:    result.FlagPresented,
					LogLevel: slog.LevelInfo,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).From(NewStr("4211", nil, false), func() (int, error) {
					return 4212, nil
				})
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    4212,
					Error:    nil,
					Flags:    result.FlagDefaulted,
					LogLevel: slog.LevelWarn,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).From(NewStr("4212", ErrTest1, true), func() (int, error) {
					return 4213, nil
				})
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest1,
					Flags:    result.FlagLookupError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).From(NewStr("4213", ErrTest2, false), func() (int, error) {
					return 4214, nil
				})
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest2,
					Flags:    result.FlagLookupError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.Int().SetConfigurator(cfg).From(NewStr("4214", nil, false), func() (int, error) {
					return 4215, ErrTest1
				})
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    0,
					Error:    ErrTest1,
					Flags:    result.FlagDefaulterError,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.
					Int().
					SetConfigurator(cfg).
					Check(func(r *result.Result[int]) (bool, string) {
						return r.Value != 4215, "Test check"
					}).
					From(NewStr("4215", nil, true))
			},
			ExpectedResults: []ExpectedResult{
				{
					Value: 4215,
					ErrorFn: func(err error) bool {
						return err != nil && err.Error() == "Test check"
					},
					Flags:    result.FlagRequired | result.FlagPresented | result.FlagCheckFailed,
					LogLevel: slog.LevelError,
				},
			},
		},
		{
			Action: func(cfg *configurator.Configurator) {
				config.
					Int().
					SetConfigurator(cfg).
					Check(func(r *result.Result[int]) (bool, string) {
						return r.Value == 4216, "Test check"
					}).
					From(NewStr("4216", nil, true))
			},
			ExpectedResults: []ExpectedResult{
				{
					Value:    4216,
					Error:    nil,
					Flags:    result.FlagRequired | result.FlagPresented,
					LogLevel: slog.LevelInfo,
				},
			},
		},
	})
}

func runConfiguratorTests[T any](t *testing.T, tests []ConfiguratorTest[T]) {
	for idx, ct := range tests {
		t.Run(fmt.Sprintf("Test %d", idx), func(t *testing.T) {
			cfg := configurator.Configurator{}

			ct.Action(&cfg)

			if expLen := len(ct.ExpectedResults); expLen != len(cfg.Results) {
				t.Fatalf(
					"Value results amount doesnt match, expected %d, got: %d\n%v\n",
					expLen,
					len(cfg.Results),
					cfg.Results,
				)
			}

			for i, expResult := range ct.ExpectedResults {
				gotResult := cfg.Results[i]

				checkValueResults[T](t, i, &expResult, gotResult)
			}
		})
	}
}

func checkValueResults[T any](
	t *testing.T,
	idx int,
	exp *ExpectedResult,
	got *result.Result[any],
) {
	if exp == nil || got == nil {
		t.Fatalf("Result %d: some value results are nil\n", idx)
	}

	if exp.Flags != got.Flags {
		t.Fatalf(
			"Result %d: flags are not equal, exp: %s (%v), got: %s (%v), err: %v\n",
			idx,
			exp.Flags,
			exp.Flags.Pairs(),
			got.Flags,
			got.Flags.Pairs(),
			got.Error,
		)
	}

	if exp.ErrorFn != nil {
		if !exp.ErrorFn(got.Error) {
			t.Fatalf(
				"Result %d: error check has failed by ErrorFn, error %v\n",
				idx,
				got.Error,
			)
		}
	} else if !errors.Is(got.Error, exp.Error) {
		t.Fatalf(
			"Result %d: errors are not equal, exp: %v, got: %v\n",
			idx,
			exp.Error,
			got.Error,
		)
	}

	if exp.Value != got.Value {
		t.Fatalf(
			"Result %d: values are not equal, exp: %v (%T), got: %v (%T)\n",
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
		hasValue = hasValue || a.Key == result.LogAttrValue
		return true
	})

	if hasValue {
		t.Fatalf(
			"Result %d: log record contains `value` even when not required",
			idx,
		)
	}

	if exp.LogLevel != 0 && exp.LogLevel != logEntry.Level {
		t.Fatalf(
			"Result %d: log record levels are not equal, exp: %v, got: %v\n",
			idx,
			exp.LogLevel,
			logEntry.Level,
		)
	}
}
