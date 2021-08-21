package plugins

import (
	"github.com/kamontat/fthelper/generator/v4/src/clusters"
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func Json(data maps.Mapper, config maps.Mapper) runners.Runner {
	return clusters.NewRunner(data, config, func(p *clusters.ExecutorParameter) error {
		templates, err := fs.Build(p.Data.Mi("inputs"), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get inputs information")
			return err
		}

		output, err := fs.Build(p.Data.Mi("output"), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		content, err := configs.LoadConfigFromFileSystem(
			templates.Multiple(),
			p.Config,
			p.Data.Mi("merger"),
		)
		if err != nil {
			p.Logger.Error("cannot load template as json")
			return err
		}

		var outfile = output.Single()
		err = outfile.Build()
		if err != nil {
			p.Logger.Error("cannot build output file")
			return err
		}

		json, err := maps.ToFormatJson(content)
		if err != nil {
			p.Logger.Error("cannot create formatted json")
			return err
		}

		return outfile.Write(json)
	}, &clusters.Settings{
		DefaultWithCluster: true,
	})
}
