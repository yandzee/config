package configurator

import (
	"log/slog"
	"time"

	"github.com/yandzee/config/pkg/source"
)

type ValueResult[T any] struct {
	Source source.StringSource
	Value  T
	Error  error
	Flags  DescriptorFlag
}

func (er *ValueResult[T]) LogAttrs() []any {
	attrs := []any{
		slog.String("name", er.Source.Name()),
		slog.String("kind", er.Source.Kind()),
	}

	attrs = append(attrs, er.Flags.LogAttrs()...)

	return attrs
}

func (er *ValueResult[T]) LogRecord(withValue bool) slog.Record {
	lvl, msg := er.LevelAndMessage()

	rec := slog.NewRecord(time.Now(), lvl, msg, 0)
	rec.Add(er.LogAttrs()...)

	if withValue {
		rec.Add(slog.Any("value", er.Value))
	}

	return rec
}

func (er *ValueResult[T]) HasIssues() bool {
	return er.Error != nil || er.IsRequiredAndNotSet() || er.Flags.IsDefaulted()
}

func (er *ValueResult[T]) IsRequiredAndNotSet() bool {
	return er.Flags.IsRequired() && !er.Flags.IsDefaulted() && !er.Flags.IsPresented()
}

func (er *ValueResult[T]) LevelAndMessage() (slog.Level, string) {
	switch {
	case er.Error != nil:
		return slog.LevelError, er.Error.Error()
	case er.IsRequiredAndNotSet():
		return slog.LevelError, "Not set"
	case er.Flags.IsDefaulted():
		return slog.LevelWarn, "Value set"
	default:
		return slog.LevelInfo, "Value set"
	}
}

func (er *ValueResult[T]) Any() *ValueResult[any] {
	return &ValueResult[any]{
		Source: er.Source,
		Value:  any(er.Value),
		Flags:  er.Flags,
		Error:  er.Error,
	}
}
