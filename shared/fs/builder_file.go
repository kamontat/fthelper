package fs

func NewFile(paths []string) (FileSystem, error) {
	return newFile(paths), nil
}
