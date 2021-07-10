package features

import (
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/loggers"
)

var global = caches.New()

func Raw(name string, deps Dependencies, executor Executor) Feature {
	return &FeatureImpl{
		name:         name,
		Dependencies: deps,

		executor: executor,

		Logger: loggers.Get("feature", name),
	}
}

func RawCache(name string, deps Dependencies, executor Executor) Feature {
	if global.Has(name) {
		return global.Get(name).Data().(*FeatureImpl)
	}

	var raw = Raw(name, deps, executor)
	global.S(name, raw, "")

	return raw
}
