package commands

import (
	"github.com/kamontat/fthelper/shared/maps"
)

func getName(commands, args []string, config maps.Mapper) (string, []string) {
	var name = config.Mi("internal").So("command", DEFAULT)
	if len(args) > 0 {
		for _, cmd := range commands {
			if args[0] == cmd {
				name = cmd
				args = args[1:]
				break
			}
		}
	}

	return name, args
}
