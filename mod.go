package config

import (
	c "github.com/yandzee/config/configurator"
	"github.com/yandzee/config/transform"
)

func Of[T any](target *T, cfgrs ...*c.Configurator) *c.Getter[T] {
	g := empty[T](cfgrs...)
	g.Target = target

	return g
}

func Int(cfgrs ...*c.Configurator) *c.Getter[int] {
	return empty[int](cfgrs...)
}

func empty[T any](cfgrs ...*c.Configurator) *c.Getter[T] {
	var configurator *c.Configurator

	for _, cfgr := range cfgrs {
		if cfgr == nil {
			continue
		}

		configurator = cfgr
		break
	}

	return &c.Getter[T]{
		Configurator: configurator,
		Transformers: []transform.Transformer{},
	}
}
