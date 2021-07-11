package plugins

import (
	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/maps"
)

// SupportEnv will create internal.environments when --env is exist
func SupportEnv(p *PluginParameter) error {
	p.NewFlags(flags.String{
		Name:    "env",
		Default: "",
		Usage:   "setup output environment",
		Action: func(data string) maps.Mapper {
			return maps.New().Set("internal.environment", data)
		},
	})

	return nil
}
