package plugins

import (
	"strings"

	"github.com/kamontat/fthelper/generator/v4/src/runner"
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

		var files = make([]fs.FileSystem, 0)
		if input.IsSingle() {
			directory, err := fs.NewDirectory(fs.Next(input.Single(), p.FsConfig.Mi("variables").Si("config")))
			if err != nil {
				p.Logger.Error("cannot get find freqtrade configs template directory")
				return err
			}
			files = []fs.FileSystem{directory}
		} else if input.IsMultiple() {
			files = input.Multiple()
		}

		content, err := configs.LoadConfigFromFileSystem(files, p.Config, p.Data.Mi("merger"))
		if err != nil {
			p.Logger.Error("cannot load template data")
			return err
		}
		json, err := maps.ToFormatJson(content)
		if err != nil {
			p.Logger.Error("cannot format config to json")
			return err
		}

		var filename strings.Builder
		filename.WriteString("config")
		if p.Data.Bo("withEnv", false) {
			filename.WriteString("-" + p.Config.Mi("internal").Si("environment"))
		}
		filename.WriteString(".json")
		output, err := fs.Build(p.Data.So("output", "freqtrade"), p.FsConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}
		file, err := fs.NewFile(fs.Next(output.Single(), p.FsConfig.Mi("variables").Si("userdata"), filename.String()))
		if err != nil {
			p.Logger.Error("cannot get find freqtrade configs directory")
			return err
		}

		err = file.Build()
		if err != nil {
			p.Logger.Error("cannot build output directory")
			return err
		}
		return file.Write(json)
	})
}
