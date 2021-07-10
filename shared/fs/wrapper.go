package fs

type wrapper struct {
	Mode Mode
	fs   []FileSystem
}

func (w *wrapper) Single() FileSystem {
	if w.Mode == MULTIPLE {
		panic("Cannot get single file-system from multiple mode")
	}

	return w.fs[0]
}

func (w *wrapper) Multiple() []FileSystem {
	if w.Mode == SINGLE {
		panic("Cannot get multiple file-system from single mode")
	}

	return w.fs
}

// All will return every file system in wrapper without validation
func (w *wrapper) All() []FileSystem {
	return w.fs
}

func newWrapper(mode Mode, fs []FileSystem) *wrapper {
	return &wrapper{
		Mode: mode,
		fs:   fs,
	}
}
