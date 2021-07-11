package plugins

import (
	"github.com/kamontat/fthelper/generator/v4/src/runner"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func Empty(data maps.Mapper, config maps.Mapper) runners.Runner {
	return runner.New(data, config, func(p *runner.ExecutorParameter) error {
		return nil
	})
}
