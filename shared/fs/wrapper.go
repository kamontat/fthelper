package fs

type wrapper struct {
	fs []FileSystem
}

func (w *wrapper) Single() FileSystem {
	return w.fs[0]
}

func (w *wrapper) Multiple() []FileSystem {
	return w.fs
}

func newWrapper(fs []FileSystem) *wrapper {
	return &wrapper{
		fs: fs,
	}
}
