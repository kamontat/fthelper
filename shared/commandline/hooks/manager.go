package hooks

import (
	"fmt"
	"strings"

	"github.com/kamontat/fthelper/shared/maps"
)

type Manager struct {
	hooks map[Type][]Hook
}

func (m *Manager) Add(t Type, hook Hook) {
	m.hooks[t] = append(m.hooks[t], hook)
}

func (m *Manager) Start(t Type, config maps.Mapper) error {
	for _, hook := range m.hooks[t] {
		var err = hook(config)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) String() string {
	var str strings.Builder
	for key, value := range m.hooks {
		str.WriteString(fmt.Sprintf("%s: %d jobs\n", key, len(value)))
	}
	return str.String()
}
