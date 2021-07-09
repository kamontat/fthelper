package fs

func NewDirectory(paths []string) (*wrapper, error) {
	return newWrapper([]FileSystem{newDirectory(paths)}), nil
}
