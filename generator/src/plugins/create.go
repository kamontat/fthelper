package plugins

import (
	"github.com/kamontat/fthelper/generator/v4/src/clusters"
	"github.com/kamontat/fthelper/shared/errors"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func Create(data maps.Mapper, config maps.Mapper) runners.Runner {
	return clusters.NewRunner(data, config, func(p *clusters.ExecutorParameter) error {
		var output, err = fs.Build(p.Data.Si("output"), p.FsConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		var errs = errors.New()
		for _, f := range output.All() {
			errs.And(f.Build())
		}

		return errs.Error()
	}, &clusters.Settings{
		DefaultWithCluster: true,
	})
}
