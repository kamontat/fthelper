package plugins

import (
	"github.com/kamontat/fthelper/generator/v4/src/generators/runner"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/xtemplates"
)

// CStrategy is custom plugins for and only for freqtrade strategy
func CStrategy(data maps.Mapper, fsConfig maps.Mapper) runners.Runner {
	return runner.New(data, fsConfig, func(p *runner.ExecutorParameter) error {
		input, err := fs.Build(p.Data.So("input", "template"), p.FsConfig)
		if err != nil {
			p.Logger.Error("cannot get input information")
			return err
		}
		strategy, err := fs.NewFile(fs.Next(input.Single(), p.FsConfig.Mi("variables").So("strategy", "strategies"), p.Data.Si("name")))
		if err != nil {
			p.Logger.Error("cannot get find strategy template directory")
			return err
		}
		template, err := xtemplates.File(strategy.Basename(), strategy.Abs())
		if err != nil {
			p.Logger.Error("cannot load template data")
			return err
		}

		freqtrade, err := fs.Build(p.Data.So("output", "freqtrade"), p.FsConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}
		output, err := fs.NewFile(fs.Next(freqtrade.Single(), p.FsConfig.Mi("variables").So("userdata", "user_data"), p.FsConfig.Mi("variables").So("strategy", "strategies"), p.Data.Si("name")))
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
	})
}
