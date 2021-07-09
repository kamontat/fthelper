package fs

func NewFile(paths []string) (*wrapper, error) {
	return newWrapper([]FileSystem{newFile(paths)}), nil
}
