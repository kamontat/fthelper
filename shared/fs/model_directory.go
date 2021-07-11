package fs

import (
	"fmt"
	"os"
	"path"
)

type directory struct {
	size   int
	paths  []string
	result Result

	abs  string
	stat *os.FileInfo
}

func (f *directory) Paths() []string {
	return append([]string{""}, f.paths...)
}

func (f *directory) Abs() string {
	if f.abs != "" {
		return f.abs
	}

	var abs = string(Separator) + path.Join(f.paths...)
	f.abs = toNormalize(abs) // update cache
	if abs != f.abs {
		f.paths = toPaths(f.abs) // if abs did some normalize, update paths as well
	}

	return f.abs
}

func (f *directory) Dirpath() string {
	return toDirpath(f.Abs())
}

func (f *directory) Dirname() string {
	return toDirname(f.paths)
}

func (f *directory) Basename() string {
	return toBasename(f.paths)
}

func (f *directory) Name() string {
	return toName(f.paths)
}

func (f *directory) Relative(input FileSystem) string {
	return toRelative(f, input)
}

func (f *directory) Type() Type {
	return DIRECTORY
}

func (f *directory) IsDir() bool {
	return true
}

func (f *directory) IsFile() bool {
	return false
}

func (f *directory) Build() error {
	err := os.MkdirAll(f.Abs(), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (f *directory) ReadDir() ([]FileSystem, error) {
	var result = make([]FileSystem, 0)
	entries, err := os.ReadDir(f.Abs())
	if err != nil {
		return result, err
	}

	for _, entry := range entries {
		name := entry.Name()
		paths := make([]string, 0)
		paths = append(paths, "")         // add empty string make path is abs
		paths = append(paths, f.paths...) // add current path
		paths = append(paths, name)       // add new filename

		if entry.IsDir() {
			dir, err := NewDirectory(paths)
			if err != nil {
				return result, err
			}

			files, err := dir.ReadDir()
			if err != nil {
				return result, err
			}

			result = append(result, files...)
		} else {
			file, err := NewFile(paths)
			if err != nil {
				return result, err
			}

			result = append(result, file)
		}
	}

	return result, nil
}

func (f *directory) Reader() (*os.File, error) {
	return nil, fmt.Errorf("cannot create reader from directory")
}

func (f *directory) Read() ([]byte, error) {
	return make([]byte, 0), fmt.Errorf("cannot read data from directory")
}

func (f *directory) Writer() (*os.File, error) {
	return nil, fmt.Errorf("cannot create writer from directory")
}

func (f *directory) Write(bytes []byte) error {
	return fmt.Errorf("cannot write data to directory")
}

func (f *directory) Chmod(mode os.FileMode) error {
	return os.Chmod(f.Abs(), mode)
}

func (f *directory) Stat() (*os.FileInfo, error) {
	if f.stat != nil {
		return f.stat, nil
	} else if f.result == MISSING_RESULT {
		return nil, os.ErrNotExist
	}

	stat, err := os.Stat(f.Abs())
	if err != nil {
		if os.IsNotExist(err) {
			f.result = MISSING_RESULT
		}

		return nil, err
	}

	f.stat = &stat
	return f.stat, nil
}

func newDirectory(paths []string) *directory {
	return &directory{
		size:   len(paths),
		paths:  paths,
		result: EMPTY_RESULT,

		abs:  "",
		stat: nil,
	}
}
