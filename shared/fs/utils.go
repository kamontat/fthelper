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

// ToFiles will resolve all directory into file
// The result of this function can be guarantee type file
func ToFiles(input []FileSystem) ([]FileSystem, error) {
	return toFiles(make([]FileSystem, 0), input)
}

func toFiles(base, files []FileSystem) ([]FileSystem, error) {
	for _, file := range files {
		stat, err := file.Stat()
		if err != nil {
			return base, err
		}

		if (*stat).IsDir() {
			directory, err := NewDirectory(file.Paths())
			if err != nil {
				return base, err
			}

			files, err := directory.ReadDir()
			if err != nil {
				return base, err
			}

			return toFiles(base, files)
		} else {
			base = append(base, file)
		}
	}

	return base, nil
}
