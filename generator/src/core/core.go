package core

import (
	"github.com/kamontat/fthelper/generator/v4/src/generators"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

type Generator struct {
	config maps.Mapper
	cache  *caches.Service
	logger *loggers.Logger
}

func (g *Generator) Start() error {
	var collection, err = generators.Parse(g.config)
	if err != nil {
		return err
	}

	var group = runners.NewGroup()
	group.New(collection).Run(loggers.IsDebug()).Log(g.logger)
	return nil
}

func New(cache *caches.Service, config maps.Mapper) *Generator {
	return &Generator{
		config: config,
		cache:  cache,
		logger: loggers.Get("generator"),
	}
}
