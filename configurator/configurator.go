package configurator

import (
	"log/slog"

	"github.com/yandzee/config/result"
)

type LogOption struct{}

var LogWithValue = LogOption{}

type Configurator struct {
	Results []*result.Result[any]
}

func (c *Configurator) LogRecords(opts ...LogOption) []slog.Record {
	recs := make([]slog.Record, len(c.Results))

	withValue := false
	for _, opt := range opts {
		withValue = withValue || opt == LogWithValue
	}

	for i, vr := range c.Results {
		recs[i] = vr.LogRecord(withValue)
	}

	return recs
}

func (c *Configurator) Clear() {
	c.Results = []*result.Result[any]{}
}
