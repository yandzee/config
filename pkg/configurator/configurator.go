package configurator

import (
	"log/slog"

	"github.com/yandzee/config/pkg/source"
)

type Configurator struct {
	LogRecords    []slog.Record
	IsValueLogged bool
}

func (c *Configurator) Env(varName string) *Getter {
	return &Getter{
		Configurator: c,
		Source: &source.EnvVarSource{
			VarName: varName,
		},
	}
}
