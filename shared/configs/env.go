package configs

import (
	"os"

	"github.com/kamontat/fthelper/shared/maps"
)

// Return map of string, client will decide how to parse string data
func LoadConfigFromEnv() (maps.Mapper, error) {
	var result = maps.New()
	for _, env := range os.Environ() {
		if k, v, ok := ParseOverride(env); ok {
			if key, ok := EnvToKey(k); ok {
				result.Set(key, v)
			}
		}
	}

	return result, nil
}
