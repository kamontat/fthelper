package clusters

import (
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

type ExecutorParameter struct {
	// Executor index, this always unique
	Index int

	// Name of the executor
	Name string

	// type of current executor
	Type string

	// whole configuration mapper
	Config maps.Mapper

	// generator data
	Data maps.Mapper

	// fs configuration
	FsConfig maps.Mapper

	// helper for logging message
	Logger *loggers.Logger
}

type Executor func(p *ExecutorParameter) error
