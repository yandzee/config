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
	return c.Source(&source.EnvVarSource{
		VarName: varName,
	})
}

func (c *Configurator) Str(str string, ok ...bool) *Getter {
	return c.Source(&source.StrSource{
		Str:       str,
		Presented: len(ok) == 0 || ok[0],
	})
}

func (c *Configurator) Source(src source.StringSource) *Getter {
	return &Getter{
		Source:        src,
		IsValueLogged: c.IsValueLogged,
		LogRecords:    &c.LogRecords,
	}
}
