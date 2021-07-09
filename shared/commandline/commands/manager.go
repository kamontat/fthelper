package commands

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/maps"
)

type Manager struct {
	keys     []string
	commands map[string]*Command
}

func (m *Manager) Add(cmd *Command) {
	m.keys = append(m.keys, cmd.Name)
	m.commands[cmd.Name] = cmd
}

func (m *Manager) Get(args []string, config maps.Mapper) (*Command, []string) {
	var name, parsed = getName(m.keys, args, config)
	var cmd = m.commands[name]
	if cmd == nil {
		return &Command{
			Name: DEFAULT,
			Executor: func(parameters *ExecutorParameter) error {
				parameters.Logger.Error("You didn't specify command name to run")
				parameters.Logger.Error("Supported commands: %v", m.keys)
				return nil
			},
		}, parsed
	}

	return cmd, parsed
}

func (m *Manager) String() string {
	return fmt.Sprintf("%v", m.keys)
}
