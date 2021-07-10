package configs

import (
	"strings"
)

func ParseOverride(str string) (key, value string, ok bool) {
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
