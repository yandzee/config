package configurator

import "log/slog"

type LogOption struct{}

var LogWithValue = LogOption{}

type Configurator struct {
	ValueResults []*ValueResult[any]
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

func (c *Configurator) Clear() {
	c.ValueResults = []*ValueResult[any]{}
}
