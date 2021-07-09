package dotenv

import (
	"os"
	"strings"

	"github.com/kamontat/fthelper/shared/fs"
)

func loadFile(file fs.FileSystem, overload bool) error {
	envMap, err := readFile(file)
	if err != nil {
		return err
	}

	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range envMap {
		if !currentEnv[key] || overload {
			os.Setenv(key, value)
		}
	}

	return nil
}

func readFile(file fs.FileSystem) (envMap map[string]string, err error) {
	reader, err := file.Reader()
	if err != nil {
		return
	}

	defer reader.Close()
	return Parse(reader)
}
