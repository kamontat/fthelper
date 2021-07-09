package hooks

func New() *Manager {
	return &Manager{
		hooks: make(map[Type][]Hook),
	}
}
