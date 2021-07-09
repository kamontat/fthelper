package fs

func NewDirectories(paths [][]string) (*wrapper, error) {
	var result = []FileSystem{}
	for _, path := range paths {
		result = append(result, newDirectory(path))
	}
	return newWrapper(result), nil
}
