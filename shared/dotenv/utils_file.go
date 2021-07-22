package dotenv

import "github.com/kamontat/fthelper/shared/fs"

func ResolveFiles(files []fs.FileSystem) ([]fs.FileSystem, error) {
	return resolveFiles(make([]fs.FileSystem, 0), files)
}

func resolveFiles(base, files []fs.FileSystem) ([]fs.FileSystem, error) {
	for _, file := range files {
		stat, err := file.Stat()
		if err != nil {
			return base, err
		}

		if (*stat).IsDir() {
			directory, err := fs.NewDirectory(file.Paths())
			if err != nil {
				return base, err
			}

			files, err := directory.ReadDir()
			if err != nil {
				return base, err
			}

			return resolveFiles(base, files)
		} else {
			base = append(base, file)
		}
	}

	return base, nil
}
