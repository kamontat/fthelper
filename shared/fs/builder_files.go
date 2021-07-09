package fs

func NewFiles(paths [][]string) (*wrapper, error) {
	var result = []FileSystem{}
	for _, path := range paths {
		result = append(result, newFile(path))
	}
	return newWrapper(result), nil
}
