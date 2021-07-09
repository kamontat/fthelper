package plugins

import (
	"os"
	"path"

	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/commandline/hooks"
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/maps"
)

// Support is example how we can implement plugins support
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
			return maps.New().Set("fs.variable.current", data)
		},
	})

	p.NewFlags(flags.Array{
		Name:    "configs",
		Default: []string{path.Join(wd, "configs")},
		Usage:   "configuration directory, must contains only json files",
		Action: func(data []string) maps.Mapper {
			return maps.New().
				Set("fs.config.type", "directory").
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
