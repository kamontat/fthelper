package fs

import (
	"os"

	"github.com/kamontat/fthelper/shared/loggers"
)

type Type string

const (
	AUTO      Type = "auto"
	FILE      Type = "file"
	DIRECTORY Type = "directory"
)

func ToType(s string) (Type, bool) {
	switch s {
	case "auto":
		return AUTO, true
	case "file", "f":
		return FILE, true
	case "directory", "dir", "d":
		return DIRECTORY, true
	}

	return AUTO, false
}

func resolveAutoType(t Type, filepath string) Type {
	if t != AUTO {
		return t
	}

	var logger = loggers.Get("fs", "type", "resolver")
	var defaultValue = FILE
	var s, e = os.Stat(filepath)
	if e != nil {
		logger.Debug("fallback to %s because error occurred (%v)", defaultValue, e)
		return defaultValue
	}

	if s.IsDir() {
		logger.Debug("mark %s as directory", filepath)
		return DIRECTORY
	} else {
		logger.Debug("mark %s as file", filepath)
		return FILE
	}
}
