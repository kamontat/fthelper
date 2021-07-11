package plugins

import (
	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/maps"
)

// SupportCluster will create internal.cluster when --cluster is exist
func SupportCluster(p *PluginParameter) error {
	p.NewFlags(flags.String{
		Name:    "cluster",
		Default: "",
		Usage:   "setup output cluster",
		Action: func(data string) maps.Mapper {
			return maps.New().Set("internal.cluster", data)
		},
	})

	return nil
}
