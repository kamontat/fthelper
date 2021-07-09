package commands

import (
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/commandline/models"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

type ExecutorParameter struct {
	Name   string
	Meta   *models.Metadata
	Config maps.Mapper
	Cache  *caches.Service
	Logger *loggers.Logger
	Args   []string
}

type Executor func(p *ExecutorParameter) error
