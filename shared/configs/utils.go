package configs

import (
	"strings"
)

func parseOverride(str string) (key, value string, ok bool) {
	ok = false
	var arr = strings.Split(str, "=")
	if len(arr) != 2 {
		return
	}

	ok = true
	key = arr[0]
	value = arr[1]
	return
}

// func parseOverrides(array []string) (maps.Mapper, error) {
// 	var errs = make([]string, 0)
// 	var output = make(maps.Mapper)
// 	for _, data := range array {
// 		var arr = strings.Split(data, "=")
// 		if len(arr) != 2 {
// 			errs = append(errs, data)
// 		}
// 		output.Set(arr[0], arr[1])
// 	}

// 	if len(errs) > 0 {
// 		return output, fmt.Errorf("cannot parse follow data [ %v ]", errs)
// 	}
// 	return output, nil
// }
