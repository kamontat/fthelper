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
	var runners, err = Parse(g.config)
	if err != nil {
		return err
	}

	runners.Run().Log(g.logger)
	return nil
}

func New(cache *caches.Service, config maps.Mapper) *Generator {
	return &Generator{
		config: config,
		cache:  cache,
		logger: loggers.Get("generator"),
	}
}
