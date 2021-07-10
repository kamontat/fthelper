package plugins

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/maps"
)

// SupportFSVar will add --fsvar "<name>=<value>" for assign data to fs.variable
func SupportFSVar(p *PluginParameter) error {
	p.NewFlags(flags.Array{
		Name:    "fsvar",
		Default: []string{},
		Usage:   "add data to fs.variable config",
		Action: func(data []string) maps.Mapper {
			var m = maps.New()
			for _, d := range data {
				var key, value, ok = configs.ParseOverride(d)
				if ok {
					m.Set(fmt.Sprintf("fs.variable.%s", key), value)
				}
			}
			return m
		},
	})
	return nil
}
