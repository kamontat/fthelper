package plugins

import (
	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/maps"
)

// SupportCluster will create internal.cluster when --cluster is exist
func SupportCluster(p *PluginParameter) error {
	p.NewFlags(flags.Array{
		Name:    "clusters",
		Default: []string{},
		Usage:   "setup output clusters",
		Action: func(data []string) maps.Mapper {
			return maps.New().Set("internal.clusters", data)
		},
	})

	return nil
}
