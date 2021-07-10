package main

import (
	"os"

	"github.com/kamontat/fthelper/generator/v4/src/core"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/commandline"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/kamontat/fthelper/shared/commandline/models"
	"github.com/kamontat/fthelper/shared/commandline/plugins"
)

var (
	name    string = "ftgenerator"
	version string = "dev"
	commit  string = "none"
	date    string = "unknown"
	builtBy string = "manually"
)

func main() {
	var cmd = commandline.New(caches.Global, &models.Metadata{
		Name:    name,
		Version: version,
		Commit:  commit,
		Date:    date,
		BuiltBy: builtBy,
	}).Plugin(plugins.SupportLogLevel).
		Plugin(plugins.SupportVersion).
		Plugin(plugins.SupportListConfig).
		Plugin(plugins.SupportFSVar).
		Plugin(plugins.SupportDotEnv). // dotenv must come before config
		Plugin(plugins.SupportConfig).
		Plugin(plugins.SupportEnv).
		Plugin(plugins.SupportBanner).
		Command(&commands.Command{
			Name: commands.DEFAULT,
			Executor: func(p *commands.ExecutorParameter) error {
				return core.New(p.Cache, p.Config).Start()
			},
		})

	var err = cmd.Start(os.Args)
	if err != nil {
		panic(err)
	}
}
