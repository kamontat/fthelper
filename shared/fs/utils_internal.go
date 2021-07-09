package fs

import (
	"path"
	"strings"
)

func toPaths(path string) []string {
	return strings.Split(path, Separator)
}

func toDirname(paths []string) string {
	return paths[len(paths)-2]
}

func toDirpath(abs string) string {
	return path.Dir(abs)
}

func toBasename(paths []string) string {
	return paths[len(paths)-1]
}

func toName(paths []string) string {
	filename := toBasename(paths)
	return strings.Replace(filename, path.Ext(filename), "", 1)
}

// find relative path of a from b
func toRelative(a, b FileSystem) string {
	return strings.Replace(a.Abs(), b.Abs(), "", 1)[1:]
}
