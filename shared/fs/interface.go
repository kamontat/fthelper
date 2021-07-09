package fs

import "os"

type FileSystem interface {
	Abs() string
	Dirpath() string
	Dirname() string
	Basename() string
	// replace basename without extension
	Name() string
	// return relative path from input 'f'
	Relative(f FileSystem) string

	Paths() []string

	Type() Type

	IsDir() bool
	IsFile() bool

	// create parent/current directory
	Build() error

	ReadDir() ([]FileSystem, error)
	Reader() (*os.File, error)
	Read() ([]byte, error)

	Writer() (*os.File, error)
	Write(bs []byte) error

	Stat() (*os.FileInfo, error)
	Chmod(mode os.FileMode) error
}
