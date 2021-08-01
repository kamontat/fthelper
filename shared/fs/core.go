package fs

import (
	"fmt"
	"path"

	"github.com/kamontat/fthelper/shared/maps"
)

// {
//   "fs": {
//     "variables": {
//       "current": "/etc",
//       "a": "/go/to/a"
//     },
//     "fappname-1": {
//       "mode": "single", // single | multiple
//       "type": "file", // auto | file | directory
//       "fullpath": "{{ .a }}/username/data.txt"
//     },
//     "dappname-1": {
//       "mode": "single", // single | multiple
//       "type": "directory", // auto | file | directory
//       "fullpath": "/tmp/username"
//     },
//     "dsappname-1": {
//       "mode": "multiple", // single | multiple
//       "type": "file", // auto | file | directory
//       "fullpath": ["/tmp/username", "/tmp/rootname"]
//     },
//     "fappname-2": {
//       "mode": "single", // single | multiple
//       "type": "directory", // auto | file | directory
//       "paths": [
//         "{{ .current }}", "hello", "root"
//       ]
//     },
//     "fappname-3": {
//       "mode": "multiple", // single | multiple
//       "type": "directory", // auto | file | directory
//       "paths": [
//         ["/etc/test", "hello", "root"],
//         ["/etc/test", "hello", "user"],
//         ["/etc/test", "hello", "myself"],
//       ]
//     },
//   }
// }

// New will extract data from mapper base on following criteria
// 1. Try to get `fullpath` if any
// 2. Try to build fullpath with `paths`
func Build(name string, fsMapper maps.Mapper) (*wrapper, error) {
	var variable = fsMapper.Mi("variables")
	var m = fsMapper.Mi(name)

	ty, ok := ToType(m.Si("type"))
	if !ok {
		return nil, fmt.Errorf("cannot get type of '%s' file-system, (%v)", name, m)
	}
	mode, ok := ToMode(m.Si("mode"))
	if !ok {
		return nil, fmt.Errorf("cannot get mode of '%s' file-system, (%v)", name, m)
	}

	switch mode {
	case SINGLE:
		var paths, err = parseSinglePaths(m, variable)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", name, err)
		}

		ty = resolveAutoType(ty, "/"+path.Join(paths...))
		switch ty {
		case FILE:
			f, e := NewFile(paths)
			return newWrapper(mode, []FileSystem{f}), e
		case DIRECTORY:
			f, e := NewDirectory(paths)
			return newWrapper(mode, []FileSystem{f}), e
		}
	case MULTIPLE:
		var paths, err = parseMultiplePaths(m, variable)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", name, err)
		}

		var result = make([]FileSystem, 0)
		for _, fpath := range paths {
			ty = resolveAutoType(ty, "/"+path.Join(fpath...))
			switch ty {
			case FILE:
				f, e := NewFile(fpath)
				if e != nil {
					return nil, e
				}
				result = append(result, f)
			case DIRECTORY:
				f, e := NewDirectory(fpath)
				if e != nil {
					return nil, e
				}
				result = append(result, f)
			}
		}

		return newWrapper(mode, result), nil
	}

	return nil, fmt.Errorf("cannot found builder of following data (type=%s, mode=%s)", ty, mode)
}

func Next(fs FileSystem, next ...string) []string {
	var extra = fs.Paths()
	for _, n := range next {
		extra = append(extra, toPaths(n)...)
	}
	return extra
}
