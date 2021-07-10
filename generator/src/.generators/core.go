package generators

import (
	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/models"
	"github.com/kamontat/fthelper/shared/runners"
)

type resolverStatus int

const (
	noLoad resolverStatus = iota
	loading
	loaded
)

func New(name string, depends features.Dependencies, executor Executor, fs ...features.Feature) runners.Runner {
	return runners.NewRunner(name, generalValidator, func(i *runners.SingleInfo) error {
		var param = &ExecutorParameter{
			Logger: loggers.Get("generator", name),
			Cache:  caches.New(),
		}

		var input = i.Input().(models.Mapper)

		// build feature data
		for _, ifeature := range toFeatures(i.Input()) {
			feature := ifeature.(features.Feature)
			err := resolver(input, param, feature, "")
			if err != nil {
				return err
			}
		}

		// Debug mode will print input information
		if input.Bo(features.MergeKeys(features.KEY_MODE, features.KEY_DEBUG), false) {
			param.Logger.Log(input)
		}

		// Disable mode will non execute anything
		if input.Bo(features.MergeKeys(features.KEY_MODE, features.KEY_DISABLE), false) {
			i.Disabled()
			return nil
		}

		return executor(param, input)
	}).Input(toMapper(depends, fs))
}
