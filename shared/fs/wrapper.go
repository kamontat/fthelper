package fs

import (
	"fmt"
	"strings"
)

type wrapper struct {
	Mode Mode
	fs   []FileSystem
}

func (w *wrapper) IsSingle() bool {
	return w.Mode == SINGLE
}

func (w *wrapper) Single() FileSystem {
	if w.Mode == MULTIPLE {
		panic("Cannot get single file-system from multiple mode")
	}

	return w.fs[0]
}

func (w *wrapper) IsMultiple() bool {
	return w.Mode == MULTIPLE
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

func (w *wrapper) String() string {
	var str strings.Builder
	for i, fs := range w.fs {
		if i >= 1 {
			str.WriteString(", ")
		}
		str.WriteString(fmt.Sprintf("%s (%s)", fs.Abs(), fs.Type()))
	}

	return fmt.Sprintf("wrapper of '%s': [%s]", w.Mode, str.String())
}

func newWrapper(mode Mode, fs []FileSystem) *wrapper {
	return &wrapper{
		Mode: mode,
		fs:   fs,
	}
}
