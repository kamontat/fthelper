package fs

func NewDirectory(paths []string) (*wrapper, error) {
	return newWrapper(SINGLE, []FileSystem{newDirectory(paths)}), nil
}
