package configs

import (
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
)

func LoadConfigFromFileSystem(files []fs.FileSystem, strategy maps.Mapper) (maps.Mapper, error) {
	var result = maps.New()
	for _, file := range files {
		if file.IsDir() {
			var files, err = file.ReadDir()
			if err != nil {
				return result, err
			}

			output, err := LoadConfigFromFileSystem(files, strategy)
			if err != nil {
				return result, err
			}

			result = maps.Merger(result).Add(output).SetConfig(strategy).Merge()
		} else {
			var content, err = file.Read()
			if err != nil {
				return result, err
			}

			output, err := maps.FromJson(content)
			if err != nil {
				return result, err
			}

			result = maps.Merger(result).Add(output).SetConfig(strategy).Merge()
		}
	}

	return result, nil
}
