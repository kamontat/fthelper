package generators

import (
	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/models"
	"github.com/kamontat/fthelper/shared/runners"
)

func Debug(name string, fs ...features.Feature) runners.Runner {
	return New(name, NoDeps, func(param *ExecutorParameter, input models.Mapper) error {
		param.Logger.Line()
		param.Logger.Log(input.Format())
		param.Logger.Line()

		return nil
	}, fs...)
}
