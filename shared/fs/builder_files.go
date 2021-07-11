package fs

func NewFiles(paths [][]string) ([]FileSystem, error) {
	var result = []FileSystem{}
	for _, path := range paths {
		result = append(result, newFile(path))
	}
	return result, nil
}
