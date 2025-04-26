package result

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/yandzee/config/source"
)

const LogAttrValue = "value"

type Result[T any] struct {
	Source source.StringSource
	Value  T
	Error  error
	Flags  ResultFlag
}

func (er *Result[T]) LogAttrs() []any {
	attrs := []any{
		slog.String("name", er.Source.Name()),
		slog.String("kind", er.Source.Kind()),
	}

	attrs = append(attrs, er.Flags.LogAttrs()...)

	return attrs
}

func (er *Result[T]) LogRecord(withValue bool) slog.Record {
	lvl, msg := er.LevelAndMessage()

	rec := slog.NewRecord(time.Now(), lvl, msg, 0)
	rec.Add(er.LogAttrs()...)

	if withValue {
		rec.Add(slog.Any(LogAttrValue, er.Value))
	}

	return rec
}

func (er *Result[T]) IsRequiredAndNotSet() bool {
	return er.Flags.IsRequired() && !er.Flags.IsDefaulted() && !er.Flags.IsPresented()
}

func (er *Result[T]) LevelAndMessage() (slog.Level, string) {
	prefix := fmt.Sprintf("%s:%s", er.Source.Kind(), er.Source.Name())

	switch {
	case er.Error != nil:
		return slog.LevelError, er.Error.Error()
	case er.IsRequiredAndNotSet():
		return slog.LevelError, fmt.Sprintf("%s is not set", prefix)
	case er.Flags.IsDefaulted():
		return slog.LevelWarn, fmt.Sprintf("%s is set", prefix)
	default:
		return slog.LevelInfo, fmt.Sprintf("%s is set", prefix)
	}
}

func (er *Result[T]) Any() *Result[any] {
	return &Result[any]{
		Source: er.Source,
		Value:  any(er.Value),
		Flags:  er.Flags,
		Error:  er.Error,
	}
}
