package generators

import (
	"fmt"

	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/runners"
)

func generalValidator(i *runners.SingleInfo) error {
	var log = loggers.Get("generator", "validate")
	var fts = toFeatures(i.Input())

	// this is dependencies of generators and features
	var dependencies = make(map[string]bool)
	// whitelist feature that implemented in generator
	dependencies[features.MergeKeys(features.KEY_MODE, features.KEY_DISABLE)] = true
	dependencies[features.MergeKeys(features.KEY_MODE, features.KEY_DEBUG)] = true

	// validate generators dependencies
	var depends = toDepends(i.Input())
	for _, depend := range depends.Keys() {
		dependencies[depend] = true
		var depType = depends[depend]
		if !fts.Has(depend) && depType.IsRequire() {
			return fmt.Errorf("generator '%s' requires feature '%s'", i.Name(), depend)
		}
	}

	// validate features dependencies
	for _, ifeature := range fts {
		var feature = ifeature.(features.Feature)
		for _, depend := range feature.Depends().Keys() {
			dependencies[depend] = true

			var depType = feature.Depends()[depend]
			if !fts.Has(depend) && depType.IsRequire() {
				return fmt.Errorf("feature '%s' requires feature '%s'", feature.Name(), depend)
			}
		}
	}

	// validate unused dependencies
	for key := range fts {
		if _, ok := dependencies[key]; !ok {
			log.Warn(fmt.Sprintf("generator '%s' not specify feature '%s' as dependency", i.Name(), key))
		}
	}

	return nil
}
