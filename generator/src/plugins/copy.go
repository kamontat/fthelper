package plugins

import (
	"github.com/kamontat/fthelper/generator/v4/src/runner"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func Copy(data maps.Mapper, fsConfig maps.Mapper) runners.Runner {
	return runner.New(data, fsConfig, func(p *runner.ExecutorParameter) error {
		input, err := fs.Build(p.Data.Si("input"), p.FsConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		output, err := fs.Build(p.Data.Si("output"), p.FsConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		return fs.Copy(input.Single(), output.Single())
	})
}
