package commands

func New() *Manager {
	return &Manager{
		keys:     make([]string, 0),
		commands: make(map[string]*Command),
	}
}
