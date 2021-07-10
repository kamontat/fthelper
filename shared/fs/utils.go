package fs

import (
	"fmt"
	"io"
)

func copyDir(a, b FileSystem) error {
	var afiles, err = a.ReadDir()
	if err != nil {
		return err
	}

	for _, afile := range afiles {
		relative := afile.Relative(a)

		out := newFile(Next(b, relative))
		return copyFile(afile, out)
	}

	return nil
}

func copyDirFiles(a []FileSystem, b FileSystem) error {
	for _, file := range a {
		var out = newFile(Next(b, file.Basename()))
		var err = copyFile(file, out)
		if err != nil {
			return err
		}
	}
	return nil
}

func copyFile(a, b FileSystem) error {
	reader, err := a.Reader()
	if err != nil {
		return err
	}

	err = b.Build()
	if err != nil {
		return err
	}
	writer, err := b.Writer()
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, reader)
	return err
}

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
