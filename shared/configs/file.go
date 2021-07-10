package configs

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtemplates"
)

func LoadConfigFromFileSystem(files []fs.FileSystem, fsVariable maps.Mapper, strategy maps.Mapper) (maps.Mapper, error) {
	var result = maps.New()
	for _, file := range files {
		if file.IsDir() {
			var files, err = file.ReadDir()
			if err != nil {
				return result, err
			}

			output, err := LoadConfigFromFileSystem(files, fsVariable, strategy)
			if err != nil {
				return result, err
			}

			result = maps.Merger(result).Add(output).SetConfig(strategy).Merge()
		} else {
			// read content
			var content, err = file.Read()
			if err != nil {
				return result, err
			}

			// compile template data
			parsedContent, err := xtemplates.Text(string(content), fsVariable)
			if err != nil {
				return result, err
			}

			// convert content to mapper
			output, err := maps.FromJson([]byte(parsedContent))
			if err != nil {
				fmt.Println(parsedContent)
				return result, err
			}

			// merge result together
			result = maps.Merger(result).Add(output).SetConfig(strategy).Merge()
		}
	}

	return result, nil
}
