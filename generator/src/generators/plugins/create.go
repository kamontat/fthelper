package plugins

import (
	"github.com/kamontat/fthelper/generator/v4/src/generators/runner"
	"github.com/kamontat/fthelper/shared/errors"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func Create(data maps.Mapper, fsConfig maps.Mapper) runners.Runner {
	return runner.New(data, fsConfig, func(p *runner.ExecutorParameter) error {
		var output, err = fs.Build(p.Data.Si("output"), p.FsConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		var errs = errors.New(p.Logger.Name)
		for _, f := range output.All() {
			errs.And(f.Build())
		}

		return errs.Error()
	})
}
