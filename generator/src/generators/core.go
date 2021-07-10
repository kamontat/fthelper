package generators

import (
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

type Generator struct {
	config maps.Mapper
	cache  *caches.Service
	logger *loggers.Logger
}

func (g *Generator) Start() error {
	return nil
}

func New(cache *caches.Service, config maps.Mapper) *Generator {
	return &Generator{
		config: config,
		cache:  cache,
		logger: loggers.Get("generator"),
	}
}
