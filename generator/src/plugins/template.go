package plugins

import (
	"github.com/kamontat/fthelper/generator/v4/src/runner"
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/xtemplates"
)

// TODO: support load files from directory
func Template(data maps.Mapper, fsConfig maps.Mapper) runners.Runner {
	return runner.New(data, fsConfig, func(p *runner.ExecutorParameter) error {
		input, err := fs.Build(p.Data.Si("input"), p.FsConfig)
		if err != nil {
			p.Logger.Error("cannot get input information")
			return err
		}

		infile := input.Single()
		template, err := xtemplates.File(infile.Basename(), infile.Abs())
		if err != nil {
			p.Logger.Error("cannot create template object")
			return err
		}

		output, err := fs.Build(p.Data.Si("output"), p.FsConfig)
		if err != nil {
			p.Logger.Error("cannot get output information")
			return err
		}

		err = output.Single().Build()
		if err != nil {
			p.Logger.Error("cannot build output directory")
			return err
		}
		writer, err := output.Single().Writer()
		if err != nil {
			p.Logger.Error("cannot get output writer")
			return err
		}

		return template.Execute(writer, p.Config)
	})
}
