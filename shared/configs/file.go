package configs

import (
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtemplates"
)

func LoadConfigFromFileSystem(input []fs.FileSystem, data maps.Mapper, strategy maps.Mapper) (maps.Mapper, error) {
	var result = maps.New()
	var files, err = fs.ToFiles(input)
	if err != nil {
		return result, err
	}

	for _, file := range files {
		// read content
		var content, err = file.Read()
		if err != nil {
			return result, err
		}

		// compile template data only if data is not empty
		// If data is empty, then no point to parse templates
		if !data.IsEmpty() {
			str, err := xtemplates.Text(string(content), data)
			if err != nil {
				return result, err
			}
			content = []byte(str)
		}

		// convert content to mapper
		output, err := maps.FromJson(content)
		if err != nil {
			return result, err
		}

		// merge result together
		result = maps.Merger(result).Add(output).SetConfig(strategy).Merge()
	}

	return result, nil
}
