package plugins

import (
	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/commandline/hooks"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

func SupportLogLevel(p *PluginParameter) error {
	p.NewFlags(flags.Int{
		Name:    "log-level",
		Default: 2,
		Usage:   "setup log level; 0 is silent and 4 is verbose",
		Action: func(data int64) maps.Mapper {
			return maps.New().Set("internal.log.level", data)
		},
	})

	p.NewFlags(flags.Bool{
		Name:    "debug",
		Default: false,
		Usage:   "mark current log to debug mode",
		Action: func(data bool) maps.Mapper {
			var m = maps.New()
			if data {
				return m.Set("internal.log.level", 4)
			}
			return m
		},
	})

	p.NewHook(hooks.BEFORE_COMMAND, func(config maps.Mapper) error {
		var level, err = config.Mi("internal").Mi("log").Ne("level")
		if err != nil {
			return err
		}

		loggers.Level(loggers.LoggerLevel(level))
		return nil
	})

	return nil
}
