package plugins

import (
	"github.com/kamontat/fthelper/generator/v4/src/clusters"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/xtemplates"
)

// CStrategy is custom plugins for and only for freqtrade strategy
func CStrategy(data maps.Mapper, config maps.Mapper) *runners.Runner {
	return clusters.NewRunnerV2(data, config, func(p *clusters.ExecutorParameter) error {
		input, err := fs.Build(fs.ToObject(p.Data.Zi("input"), p.Config), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get input information")
			return err
		}
		strategy, err := fs.NewFile(fs.Next(input.Single(), p.VarConfig.Si("strategy"), p.Data.Si("name")))
		if err != nil {
			p.Logger.Error("cannot get find strategy template directory")
			return err
		}
		template, err := xtemplates.File(strategy.Basename(), strategy.Abs())
		if err != nil {
			p.Logger.Error("cannot load template data")
			return err
		}

		freqtrade, err := fs.Build(fs.ToObject(p.Data.Zi("output"), p.Config), p.VarConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}
		output, err := fs.NewFile(fs.Next(freqtrade.Single(), p.VarConfig.Si("userdata"), p.VarConfig.Si("strategy"), p.Data.Si("name")))
		if err != nil {
			p.Logger.Error("cannot get find strategy output directory")
			return err
		}
		err = output.Build()
		if err != nil {
			p.Logger.Error("cannot create output file")
			return err
		}
		writer, err := output.Writer()
		if err != nil {
			p.Logger.Error("cannot get output writer")
			return err
		}

		return template.Execute(writer, p.Config)
	}, &clusters.Settings{
		DefaultWithCluster: true,
	})
}
