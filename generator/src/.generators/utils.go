package generators

import (
	"fmt"

	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/models"
	"github.com/kamontat/fthelper/shared/runners"
)

func resolver(input models.Mapper, param *ExecutorParameter, feature features.Feature, caller string) error {
	var key = feature.Name()
	var statusKey = key + "__status__"
	var errorKey = key + "__key__"

	var fs = toFeatures(input)

	status := noLoad
	if param.Cache.Has(statusKey) {
		status = param.Cache.Get(statusKey).Data().(resolverStatus)
	}

	// param.Logger.Debug("resolving %v (%s)", key, caller)
	if status == loaded {
		var err error = nil
		if param.Cache.Has(errorKey) {
			err = param.Cache.Get(errorKey).Data().(error)
		}

		// param.Logger.Debug("cached feature '%v' (error is '%v')", key, err)
		return err
	} else if status == loading {
		return fmt.Errorf("recursive dependencies at [ '%s', '%s' ]", feature.Name(), caller)
	}

	param.Cache.U(statusKey, loading, caches.Persistent)
	for key, depType := range feature.Depends() {
		next, ok := fs.A(key)
		if ok {
			err := resolver(input, param, next.(features.Feature), feature.Name())
			if err != nil {
				return err
			}
		} else if depType.IsRequire() {
			return fmt.Errorf("missing require feature '%s'", key)
		}
	}

	param.Cache.U(statusKey, loaded, caches.Persistent) // update status to loaded
	param.Logger.Debug("building feature '%s' (depends = %v)", feature.Name(), feature.Depends())

	var runner = feature.Run(input)
	err := runner.Validate()
	if err != nil {
		param.Cache.U(errorKey, err, caches.Persistent)
		return err
	}

	err = runner.Run()
	if err != nil {
		param.Cache.U(errorKey, err, caches.Persistent)
		return err
	}

	var single = runner.Information().(*runners.SingleInfo)
	input.Set(feature.Name(), single.Output()) // update feature output to generator input
	return nil
}
