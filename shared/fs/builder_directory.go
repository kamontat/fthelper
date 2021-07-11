package fs

func NewDirectory(paths []string) (FileSystem, error) {
	return newDirectory(paths), nil
}
