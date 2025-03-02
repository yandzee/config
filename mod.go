package config

import (
	c "github.com/yandzee/config/configurator"
	"github.com/yandzee/config/transform"
)

func Of[T any](target *T, cfgrs ...*c.Configurator) *c.Getter[T] {
	var configurator *c.Configurator

	for _, cfgr := range cfgrs {
		if cfgr == nil {
			continue
		}

		configurator = cfgr
		break
	}

	trs := []transform.Transformer{}

	return &c.Getter[T]{
		Target:       target,
		Configurator: configurator,
		Transformers: trs,
	}
}

func Int() *c.Getter[int] {
	return &c.Getter[int]{}
}
