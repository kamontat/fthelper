package plugins

import (
	"github.com/kamontat/fthelper/generator/v4/src/runner"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/xtemplates"
)

// CDocker is custom plugins for and only for freqtrade docker-compose
func CDockerCompose(data maps.Mapper, config maps.Mapper) runners.Runner {
	return runner.New(data, config, func(p *runner.ExecutorParameter) error {
		input, err := fs.Build(p.Data.So("input", "template"), p.FsConfig)
		if err != nil {
			p.Logger.Error("cannot get input information")
			return err
		}
		dcompose, err := fs.NewFile(fs.Next(input.Single(), p.FsConfig.Mi("variables").Si("docker"), p.Data.So("name", "docker-compose.yml")))
		if err != nil {
			p.Logger.Error("cannot get find docker-compose template file")
			return err
		}
		template, err := xtemplates.File(dcompose.Basename(), dcompose.Abs())
		if err != nil {
			p.Logger.Error("cannot load template data")
			return err
		}

		// validate parameters
		if !p.Config.Has("data.freqtrades") {
			p.Config.Set("data.freqtrades", []string{""})
		}

		// write to output
		freqtrade, err := fs.Build(p.Data.So("output", "freqtrade"), p.FsConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}
		output, err := fs.NewFile(fs.Next(freqtrade.Single(), "docker-compose.yml"))
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
