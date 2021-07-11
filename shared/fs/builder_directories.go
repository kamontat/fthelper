package fs

func NewDirectories(paths [][]string) ([]FileSystem, error) {
	var result = []FileSystem{}
	for _, path := range paths {
		result = append(result, newDirectory(path))
	}
	return result, nil
}
