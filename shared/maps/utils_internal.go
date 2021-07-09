package maps

import (
	"fmt"
	"strings"
)

func missError(m map[string]interface{}, key string) error {
	return fmt.Errorf("map (%v) missing key='%s'", m, key)
}

func convertError(data interface{}, datatype string) error {
	return fmt.Errorf("cannot convert %v to %s", data, datatype)
}

func getByKey(m map[string]interface{}, keys []string) (interface{}, error) {
	var value, ok = m[keys[0]]
	if !ok {
		return nil, fmt.Errorf("map (%v) missing key='%s'", m, keys[0])
	}
	if len(keys) == 1 {
		return value, nil
	}

	var nextKey = keys[1:]
	next, ok := ToMapper(value)
	if !ok {
		return nil, fmt.Errorf("value of key=%s is not map but keys is not end (%v)", keys[0], nextKey)
	}

	return getByKey(next, nextKey)
}

func forEach(m map[string]interface{}, keys []string, fn ForEachFn) {
	var result = make(map[string]interface{}) // interface{} never be map
	var update = func(key string, value interface{}) {
		result[key] = value
	}

	for k, v := range m {
		var copyKeys = append(keys, k)
		if m, ok := ToMapper(v); ok {
			forEach(m, copyKeys, update)
		} else {
			var k = strings.Join(copyKeys, ".")
			result[k] = v
		}
	}

	for key, value := range result {
		fn(key, value)
	}
}
