package datatype

import (
	"strconv"
)

// convert string to different datatype base on input type
func ConvertStringTo(s string, t interface{}) interface{} {
	switch t.(type) {
	case int, int64:
		i, e := strconv.ParseInt(s, 10, 64)
		if e == nil {
			return i
		}
	case float32, float64:
		i, e := strconv.ParseFloat(s, 64)
		if e == nil {
			return i
		}
	case bool:
		i, e := strconv.ParseBool(s)
		if e == nil {
			return i
		}
	case []interface{}, []string:
		a, ok := ForceArray(s)
		if ok {
			return a
		}
	}

	return s
}

// convert string to different datatype
func ConvertString(s string) interface{} {
	i, e := strconv.ParseInt(s, 10, 64)
	if e == nil {
		return i
	}

	f, e := strconv.ParseFloat(s, 64)
	if e == nil {
		return f
	}

	b, e := strconv.ParseBool(s)
	if e == nil {
		return b
	}

	return s
}
