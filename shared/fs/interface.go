package fs

import "os"

type FileSystem interface {
	// full absolute path from root directory
	Abs() string
	// directory path excluded current position
	Dirpath() string
	// upper directory name
	Dirname() string
	// filename with extension
	Basename() string
	// replace basename without extension
	Name() string
	// return relative path from input 'f'
	Relative(f FileSystem) string
	// full path separate by path separator
	Paths() []string

	Type() Type

	IsDir() bool
	IsFile() bool

	// create parent/current directory
	Build() error
	// read files/directories from current position
	ReadDir() ([]FileSystem, error)

	// create file reader
	Reader() (*os.File, error)
	// read content from current file path
	Read() ([]byte, error)
	// create file writer
	Writer() (*os.File, error)
	// write content to current file path
	Write(bs []byte) error

	// read file/directory information
	Stat() (*os.FileInfo, error)
	// modify permission of current file/directory
	Chmod(mode os.FileMode) error
}
