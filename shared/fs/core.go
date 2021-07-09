package fs

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/maps"
)

// {
//   "fs": {
//     "variable": {
//       "current": "/etc",
//       "a": "/go/to/a"
//     },
//     "fappname-1": {
//       "mode": "single", // single | multiple
//       "type": "file", // file | directory
//       "fullpath": "{{ .a }}/username/data.txt"
//     },
//     "dappname-1": {
//       "mode": "single", // single | multiple
//       "type": "directory", // file | directory
//       "fullpath": "/tmp/username"
//     },
//     "dsappname-1": {
//       "mode": "multiple", // single | multiple
//       "type": "file", // file | directory
//       "fullpath": ["/tmp/username", "/tmp/rootname"]
//     },
//     "fappname-2": {
//       "mode": "single", // single | multiple
//       "type": "directory", // file | directory
//       "paths": [
//         "{{ .current }}", "hello", "root"
//       ]
//     },
//     "fappname-3": {
//       "mode": "multiple", // single | multiple
//       "type": "directory", // file | directory
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
	var variable = fsMapper.Mi("variable")
	var m = fsMapper.Mi(name)

	ty, ok := ToType(m.Si("type"))
	if !ok {
		return nil, fmt.Errorf("cannot get type of file-system, (%v)", m)
	}
	mode, ok := ToMode(m.Si("mode"))
	if !ok {
		return nil, fmt.Errorf("cannot get mode of file-system, (%v)", m)
	}

	switch mode {
	case SINGLE:
		var paths, err = parseSinglePaths(m, variable)
		if err != nil {
			return nil, err
		}

		switch ty {
		case FILE:
			return NewFile(paths)
		case DIRECTORY:
			return NewDirectory(paths)
		}
	case MULTIPLE:
		var paths, err = parseMultiplePaths(m, variable)
		if err != nil {
			return nil, err
		}

		switch ty {
		case FILE:
			return NewFiles(paths)
		case DIRECTORY:
			return NewDirectories(paths)
		}
	}

	return nil, fmt.Errorf("cannot found builder of following data (type=%s, mode=%s)", ty, mode)
}
