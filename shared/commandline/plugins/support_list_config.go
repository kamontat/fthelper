package plugins

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/kamontat/fthelper/shared/commandline/flags"
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

func SupportListConfig(p *PluginParameter) error {
	p.NewCommand(&commands.Command{
		Name: "config",
		Flags: flags.New(flags.Bool{
			Name:    "data",
			Default: false,
			Usage:   "show config value as well",
			Action: func(data bool) maps.Mapper {
				return maps.New().Set("internal.flag.data", data)
			},
		}, flags.Bool{
			Name:    "all",
			Default: false,
			Usage:   "show all configuration, including internal",
			Action: func(data bool) maps.Mapper {
				return maps.New().Set("internal.flag.all", data)
			},
		}),
		Executor: func(p *commands.ExecutorParameter) error {
			var withData = p.Config.Mi("internal").Mi("flag").Bo("data", false)
			var all = p.Config.Mi("internal").Mi("flag").Bo("all", false)

			var headers = []string{"Key", "Environment"}
			if withData {
				headers = append(headers, "Type", "Value")
			}

			var table = loggers.Get().Table(uint(len(headers)))
			table.Header(headers...)

			var keys = p.Config.Keys()
			sort.Strings(keys)

			// Sorted keys
			for _, key := range keys {
				if !all && strings.HasPrefix(key, "internal") {
					continue
				}

				var env = configs.KeyToEnv(key)
				var row = []string{key, env}
				if withData {
					var value, _ = p.Config.Get(key)
					var t = fmt.Sprintf("%T", value)
					row = append(row, t, fmt.Sprintf("%v", value))
				}

				table.Row(row...)
			}

			return table.End()
		},
	})

	return nil
}
