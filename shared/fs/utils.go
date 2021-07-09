package fs

import (
	"fmt"
	"path"

	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/utils"
	"github.com/kamontat/fthelper/shared/xtemplates"
)

func parseSinglePaths(m, variable maps.Mapper) ([]string, error) {
	var paths = make([]string, 0)
	if fullpath, ok := m.S("fullpath"); ok && fullpath != "" {
		var data, err = xtemplates.Text(fullpath, variable)
		if err != nil {
			return paths, err
		}

		paths = toPaths(data)
	} else if arr, ok := m.A("paths"); ok {
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
	}

	if len(paths) < 1 {
		return paths, fmt.Errorf("cannot found path from input map (%v)", m)
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
