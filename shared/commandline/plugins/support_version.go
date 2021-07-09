package plugins

import (
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/maps"
)

func SupportVersion(p *PluginParameter) error {
	p.NewFlags(flags.Bool{
		Name:    "version",
		Default: false,
		Usage:   "show current application version",
		Action: func(data bool) maps.Mapper {
			var m = maps.New()
			if data {
				return m.Set("internal.command", "version")
			}
			return m
		},
	})
	p.NewCommand(&commands.Command{
		Name: "version",
		Executor: func(p *commands.ExecutorParameter) error {
			p.Logger.Log(p.Meta.String())
			return nil
		},
	})

	return nil
}
