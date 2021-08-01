package plugins

import (
	"os"
	"path"

	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/commandline/hooks"
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/maps"
)

// SupportConfig will load configuration from configs directory
// It also add --pwd for current directory, --config-dirs for custom config directory
func SupportConfig(p *PluginParameter) error {
	var wd, err = os.Getwd()
	if err != nil {
		return err
	}

	p.NewFlags(flags.String{
		Name:    "pwd",
		Default: wd,
		Usage:   "current directory",
		Action: func(data string) maps.Mapper {
			return maps.New().Set("fs.variables.current", data)
		},
	})

	// @deprecated - will remove on version 5.0.0
	p.NewFlags(flags.Array{
		Name:    "config-dirs",
		Default: make([]string, 0),
		Usage:   "configuration directory, must contains only json files",
		Action: func(data []string) maps.Mapper {
			if len(data) > 0 {
				p.Logger.Warn("--config-dirs is deprecated because config is now support both file and directory. please use --configs instead")
				return maps.New().
					Set("fs.config.type", "directory").
					Set("fs.config.mode", "multiple").
					Set("fs.config.fullpath", data)
			}

			return maps.New()
		},
	})

	p.NewFlags(flags.Array{
		Name:    "configs",
		Default: []string{path.Join(wd, "configs")},
		Usage:   "configuration file/directory. directory must contains only json files and file must be json",
		Action: func(data []string) maps.Mapper {
			return maps.New().
				Set("fs.config.type", "auto").
				Set("fs.config.mode", "multiple").
				Set("fs.config.fullpath", data)
		},
	})

	p.NewHook(hooks.BEFORE_COMMAND, func(config maps.Mapper) error {
		var addition, err = configs.New("config", config).Build()
		if err != nil {
			return err
		}

		addition.ForEach(func(key string, value interface{}) {
			config.Set(key, value)
		})
		return nil
	})

	return nil
}
