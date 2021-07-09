package configs

import (
	"os"

	"github.com/kamontat/fthelper/shared/maps"
)

// Return map of string, client will decide how to parse string data
func LoadConfigFromEnv(base maps.Mapper) (maps.Mapper, error) {
	var result = maps.New()
	for _, env := range os.Environ() {
		if k, v, ok := parseOverride(env); ok {
			if key, ok := EnvToKey(k); ok {
				result.Set(key, v)
			}
		}
	}

	return result, nil
}

// func EnvString(name string, def string) string {
// 	var v, ok = os.LookupEnv(name)
// 	if ok {
// 		// log.Debug("found env name '%s' = %s", name, utils.MaskString(v, utils.MEDIUM))
// 		return v
// 	}

// 	return def
// }

// func EnvInt(name string, def int64) int64 {
// 	var v, ok = os.LookupEnv(name)
// 	if !ok {
// 		return def
// 	}

// 	var i, err = strconv.ParseInt(v, 10, 64)
// 	if err != nil {
// 		return def
// 	}

// 	// log.Debug("found env name '%s' = %d", name, i)
// 	return i
// }

// func EnvFloat(name string, def float64) float64 {
// 	var v, ok = os.LookupEnv(name)
// 	if !ok {
// 		return def
// 	}

// 	var i, err = strconv.ParseFloat(v, 64)
// 	if err != nil {
// 		return def
// 	}

// 	// log.Debug("found env name '%s' = %f", name, i)
// 	return i
// }

// func EnvBool(name string, def bool) bool {
// 	var v, ok = os.LookupEnv(name)
// 	if !ok {
// 		return def
// 	}

// 	var i, err = strconv.ParseBool(v)
// 	if err != nil {
// 		return def
// 	}

// 	// log.Debug("found env name '%s' = %t", name, i)
// 	return i
// }

// func GetEnv(name string, def interface{}) interface{} {
// 	switch v := def.(type) {
// 	case int:
// 		return EnvInt(name, int64(v))
// 	case int64:
// 		return EnvInt(name, v)
// 	case float32:
// 		return EnvFloat(name, float64(v))
// 	case float64:
// 		return EnvFloat(name, v)
// 	case bool:
// 		return EnvBool(name, v)
// 	case string:
// 		return EnvString(name, v)
// 	default:
// 		// log.Debug("cannot get env value from not support type (" + name + ")")
// 		return def
// 	}
// }
