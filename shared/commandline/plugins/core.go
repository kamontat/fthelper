package plugins

func New() *Manager {
	return &Manager{
		plugins: make([]Plugin, 0),
	}
}
