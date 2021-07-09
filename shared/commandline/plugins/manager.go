package plugins

import "fmt"

type Manager struct {
	plugins []Plugin
}

func (m *Manager) Add(plugin Plugin) *Manager {
	m.plugins = append(m.plugins, plugin)
	return m
}

func (m *Manager) Build(parameters *PluginParameter) error {
	for _, plugin := range m.plugins {
		var err = plugin(parameters)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) String() string {
	return fmt.Sprintf("Manager %d plugins", len(m.plugins))
}
