package datatype

import (
	"strings"
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

// ToArray will try to convert interface{} to array of interface{}
func ToArray(a interface{}) ([]interface{}, bool) {
	adata, ok := a.([]interface{})
	if ok {
		return adata, ok
	}

	// support string array
	sdata, ok := a.([]string)
	if ok {
		s := make([]interface{}, len(sdata))
		for i, v := range sdata {
			s[i] = v
		}

		return s, ok
	}

	// support int array
	idata, ok := a.([]int)
	if ok {
		s := make([]interface{}, len(idata))
		for i, v := range idata {
			s[i] = v
		}

		return s, ok
	}

	// support string with `a:1,2,3` format
	if s, ok := a.(string); ok {
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

// ForceArray will force to convert interface{} to array of interface{}
func ForceArray(i interface{}) []interface{} {
	if a, ok := ToArray(i); ok {
		return a
	}

	if s, ok := i.(string); ok {
		v, _ := parseArray(s)
		var t = make([]interface{}, 0)
		for _, d := range v {
			t = append(t, d)
		}
		return t
	}

	// return 1 element array
	return []interface{}{i}
}
