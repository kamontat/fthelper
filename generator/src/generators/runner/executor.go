package runner

import (
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

type ExecutorParameter struct {
	Name     string
	Config   maps.Mapper // whole configuration mapper
	Data     maps.Mapper // generator data
	FsConfig maps.Mapper // fs configuration

	Logger *loggers.Logger
}

type Executor func(p *ExecutorParameter) error
