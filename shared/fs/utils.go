package fs

import (
	"fmt"
)

func Copy(a, b FileSystem) error {
	if a.IsDir() && b.IsDir() {
		return copyDir(a, b)
	} else if a.IsFile() && b.IsFile() {
		return copyFile(a, b)
	} else if a.IsFile() && b.IsDir() {
		return copyDirFiles([]FileSystem{a}, b)
	}

	return fmt.Errorf("cannot copy from directory (%s) to file (%s)", a.Abs(), b.Abs())
}
