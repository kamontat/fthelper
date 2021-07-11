package plugins

import (
	"fmt"

	"github.com/kamontat/fthelper/generator/v4/src/generators/runner"
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

// CConfig is custom plugins for and only for freqtrade config
func CConfig(data maps.Mapper, fsConfig maps.Mapper) runners.Runner {
	return runner.New(data, fsConfig, func(p *runner.ExecutorParameter) error {
		input, err := fs.Build(p.Data.So("input", "template"), p.FsConfig)
		if err != nil {
			p.Logger.Error("cannot get input information")
			return err
		}
		config, err := fs.NewDirectory(fs.Next(input.Single(), p.FsConfig.Mi("variables").So("config", "configs")))
		if err != nil {
			p.Logger.Error("cannot get find freqtrade configs template directory")
			return err
		}
		template, err := configs.LoadConfigFromFileSystem([]fs.FileSystem{config}, p.Config, p.Data.Mi("merger"))
		if err != nil {
			p.Logger.Error("cannot load template data")
			return err
		}

		fmt.Print(template)
		return nil
	})
}
