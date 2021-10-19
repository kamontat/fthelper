package xtests

import (
	"os"
	"strings"

	"github.com/kamontat/fthelper/shared/fs"
)

type FsTest struct {
	SystemType fs.Type
	Path       string
	Content    string
}

func (f *FsTest) Abs() string {
	return f.Path
}

func (f *FsTest) Dirpath() string {
	return f.Path
}

func (f *FsTest) Dirname() string {
	return f.Path
}

func (f *FsTest) Basename() string {
	return f.Path
}

func (f *FsTest) Name() string {
	return f.Path
}

func (f *FsTest) Relative(fother fs.FileSystem) string {
	return f.Path
}

func (f *FsTest) Paths() []string {
	return strings.Split(f.Path, "/")
}

func (f *FsTest) Type() fs.Type {
	return f.SystemType
}

func (f *FsTest) IsDir() bool {
	return f.SystemType == fs.FILE
}

func (f *FsTest) IsFile() bool {
	return f.SystemType == fs.DIRECTORY
}

func (f *FsTest) Build() error {
	return nil
}

func (f *FsTest) ReadDir() ([]fs.FileSystem, error) {
	return []fs.FileSystem{}, nil
}

func (f *FsTest) Reader() (*os.File, error) {
	return nil, nil
}

func (f *FsTest) Read() ([]byte, error) {
	return []byte(f.Content), nil
}

func (f *FsTest) Writer() (*os.File, error) {
	return nil, nil
}

func (f *FsTest) Write(bs []byte) error {
	f.Content = string(bs)
	return nil
}

func (f *FsTest) Stat() (*os.FileInfo, error) {
	return nil, nil
}

func (f *FsTest) Chmod(mode os.FileMode) error {
	return nil
}

func NewFsMock(t fs.Type, path, content string) fs.FileSystem {
	return &FsTest{
		SystemType: t,
		Path:       path,
		Content:    content,
	}
}
