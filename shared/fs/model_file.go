package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/kamontat/fthelper/shared/utils"
)

type file struct {
	size   int
	paths  []string
	result Result

	abs     string
	content []byte
	stat    *os.FileInfo
}

func (f *file) Paths() []string {
	return utils.CloneStringArray(f.paths)
}

func (f *file) Abs() string {
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

func (f *file) Dirpath() string {
	return toDirpath(f.Abs())
}

func (f *file) Dirname() string {
	return toDirname(f.paths)
}

func (f *file) Basename() string {
	return toBasename(f.paths)
}

func (f *file) Name() string {
	return toName(f.paths)
}

func (f *file) Relative(input FileSystem) string {
	return toRelative(f, input)
}

func (f *file) Type() Type {
	return FILE
}

func (f *file) IsDir() bool {
	return false
}

func (f *file) IsFile() bool {
	return true
}

func (f *file) Build() error {
	err := os.MkdirAll(f.Dirpath(), os.ModePerm)
	if err != nil {
		return err
	}

	// create empty file (touch)
	_, err = f.Writer()
	return err
}

func (f *file) ReadDir() ([]FileSystem, error) {
	return make([]FileSystem, 0), fmt.Errorf("cannot read directory from file")
}

func (f *file) Reader() (*os.File, error) {
	return os.Open(f.Abs())
}

func (f *file) Read() ([]byte, error) {
	if len(f.content) > 0 {
		return f.content, nil
	}

	reader, err := f.Reader()
	if err != nil {
		return f.content, err
	}
	defer reader.Close()

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return f.content, err
	}

	f.content = bytes // save data to cache
	return f.content, nil
}

func (f *file) Writer() (*os.File, error) {
	return os.OpenFile(f.Abs(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}

func (f *file) Write(bytes []byte) error {
	writer, err := f.Writer()
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = writer.Write(bytes)
	if err != nil {
		return err
	}

	return writer.Sync()
}

func (f *file) Chmod(mode os.FileMode) error {
	return os.Chmod(f.Abs(), mode)
}

func (f *file) Stat() (*os.FileInfo, error) {
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

func newFile(paths []string) *file {
	return &file{
		size:   len(paths),
		paths:  paths,
		result: EMPTY_RESULT,

		abs:     "",
		content: make([]byte, 0),
		stat:    nil,
	}
}
