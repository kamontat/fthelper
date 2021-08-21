package plugins

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/maps"
)

// SupportFSVar will add --fsvar "<name>=<value>" for assign data to fs.variables
// @deprecated SupportFSVar, use SupportVar instead
func SupportVar(p *PluginParameter) error {
	p.NewFlags(flags.Array{
		Name:    "var",
		Default: []string{},
		Usage:   "add data to variables config",
		Action: func(data []string) maps.Mapper {
			var m = maps.New()
			for _, d := range data {
				var key, value, ok = configs.ParseOverride(d)
				if ok {
					m.Set(fmt.Sprintf("variables.%s", key), value)
				}
			}
			return m
		},
	})
	return nil
}
