package generators

import (
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func Parse(config maps.Mapper) (*runners.Runners, error) {
	var log = loggers.Get("generator", "parser")

	var rs = runners.New()
	for _, i := range config.Ai("generators") {
		var mapper, ok = maps.ToMapper(i)
		if !ok {
			log.Warn("generator %v is not map", i)
		}

		var runner, err = GetRunner(mapper, config)
		if err != nil {
			return rs, err
		}

		rs.Add(runner)
	}

	return rs, nil
}
