package fs

import (
	"fmt"
	"path"
	"strings"

	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/utils"
	"github.com/kamontat/fthelper/shared/xtemplates"
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

func parseSinglePaths(m, variable maps.Mapper) ([]string, error) {
	var paths = make([]string, 0)
	if m.Has("fullpath") {
		if fullpath, ok := m.S("fullpath"); ok && fullpath != "" {
			var data, err = xtemplates.Text(fullpath, variable)
			if err != nil {
				return paths, err
			}

			paths = toPaths(data)
		} else {
			var raw, _ = m.Get("fullpath")
			return paths, fmt.Errorf("we expected fullpath to be 'string' not '%T'", raw)
		}
	} else if m.Has("paths") {
		if arr, ok := m.A("paths"); ok {
			var strings = make([]string, 0)
			for _, tpl := range arr {
				var data, err = xtemplates.Text(tpl.(string), variable)
				if err != nil {
					return paths, err
				}

				strings = append(strings, data)
			}
			var p = path.Join(strings...)
			paths = toPaths(p)
		} else {
			var raw, _ = m.Get("paths")
			return paths, fmt.Errorf("we expected paths to be '[]string' not '%T'", raw)
		}
	}

	if len(paths) < 1 {
		return paths, fmt.Errorf("either fullpath or paths key must be exist at %v", m)
	}
	return paths, nil
}

func parseMultiplePaths(m, variable maps.Mapper) ([][]string, error) {
	var paths = make([][]string, 0)
	if arr, ok := m.A("fullpath"); ok {
		for _, fullpath := range arr {
			var data, err = xtemplates.Text(fullpath.(string), variable)
			if err != nil {
				return paths, err
			}

			paths = append(paths, toPaths(data))
		}
	} else if raws, ok := m.A("paths"); ok {
		// raw should be [][]string
		for _, raw := range raws {
			if arr, ok := utils.ToArray(raw); ok {
				var strings = make([]string, 0)
				for _, tpl := range arr {
					var data, err = xtemplates.Text(tpl.(string), variable)
					if err != nil {
						return paths, err
					}

					strings = append(strings, data)
				}

				var p = path.Join(strings...)
				paths = append(paths, toPaths(p))
			}
		}
	}

	if len(paths) < 1 {
		return paths, fmt.Errorf("cannot found paths from input map (%v)", m)
	}
	return paths, nil
}
