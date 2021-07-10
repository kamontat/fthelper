package generators

import (
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func Parse(config maps.Mapper) (*runners.Collection, error) {
	var log = loggers.Get("generator", "parser")
	var collection = runners.NewCollection("default")
	for _, i := range config.Ai("generators") {
		var mapper, ok = maps.ToMapper(i)
		if !ok {
			log.Warn("generator %v is not map", i)
		}

		var runner, err = GetRunner(mapper, config)
		if err != nil {
			return collection, err
		}

		collection.Add(runner)
	}

	return collection, nil
}
