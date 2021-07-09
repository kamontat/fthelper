package configs

import (
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

func New(name string, config maps.Mapper) *Builder {
	return &Builder{
		name:     name,
		config:   config,
		override: maps.New(),
		strategy: maps.New(),

		logger: loggers.Get("config", "builder"),
	}
}
