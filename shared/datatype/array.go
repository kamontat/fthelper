package datatype

import (
	"strings"

	"github.com/kamontat/fthelper/shared/utils"
)

func parseArray(s string) (out []string, isArray bool) {
	if strings.HasPrefix(s, "a:") {
		s = strings.TrimPrefix(s, "a:")
		isArray = true
	} else if strings.HasPrefix(s, "array:") {
		s = strings.TrimPrefix(s, "array:")
		isArray = true
	}

	out = strings.Split(s, ",")
	return
}

// TODO: move ToArray from utils to datatype
func ToArray(i interface{}) ([]interface{}, bool) {
	if a, ok := utils.ToArray(i); ok {
		return a, ok
	}

	if s, ok := i.(string); ok {
		if v, ok := parseArray(s); ok {
			var t = make([]interface{}, 0)
			for _, d := range v {
				t = append(t, d)
			}
			return t, ok
		}
	}

	return make([]interface{}, 0), false
}

func ForceArray(i interface{}) ([]interface{}, bool) {
	if a, ok := utils.ToArray(i); ok {
		return a, ok
	}

	if s, ok := i.(string); ok {
		v, _ := parseArray(s)
		var t = make([]interface{}, 0)
		for _, d := range v {
			t = append(t, d)
		}
		return t, ok
	}

	return make([]interface{}, 0), false
}
