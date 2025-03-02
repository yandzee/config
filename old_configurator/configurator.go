package configurator

import (
	"log/slog"

	"github.com/yandzee/config/source"
)

type LogOption struct{}

var LogWithValue = LogOption{}

type Configurator struct {
	ValueResults []*ValueResult[any]
}

func (c *Configurator) Env(varName string) *Getter {
	return c.Source(&source.EnvVarSource{
		VarName: varName,
	})
}

func (c *Configurator) Source(src source.StringSource) *Getter {
	return &Getter{
		Source:       src,
		ValueResults: &c.ValueResults,
	}
}

func (c *Configurator) LogRecords(opts ...LogOption) []slog.Record {
	recs := make([]slog.Record, len(c.ValueResults))

	withValue := false
	for _, opt := range opts {
		withValue = withValue || opt == LogWithValue
	}

	for i, vr := range c.ValueResults {
		recs[i] = vr.LogRecord(withValue)
	}

	return recs
}
