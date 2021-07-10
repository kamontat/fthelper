package fs

func NewFile(paths []string) (*wrapper, error) {
	return newWrapper(SINGLE, []FileSystem{newFile(paths)}), nil
}
