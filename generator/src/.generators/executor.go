package generators

import (
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/models"
)

type ExecutorParameter struct {
	Logger *loggers.Logger
	Cache  *caches.Service
}

type Executor func(param *ExecutorParameter, input models.Mapper) error
