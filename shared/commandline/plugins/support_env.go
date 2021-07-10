package plugins

import (
	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/maps"
)

// SupportEnv will create internal.environments when --env is exist
func SupportEnv(p *PluginParameter) error {
	p.NewFlags(flags.Array{
		Name:    "envs",
		Default: []string{},
		Usage:   "setup output environment name",
		Action: func(data []string) maps.Mapper {
			return maps.New().Set("internal.environments", data)
		},
	})

	return nil
}
