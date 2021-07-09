package flags

func New(flags ...Flag) *Manager {
	var m = make(map[string]Flag)
	for _, flag := range flags {
		m[flag.FlagName()] = flag
	}

	return &Manager{
		keys:  make([]string, 0),
		flags: m,
	}
}
