package datatype

import (
	"strconv"
)

// convert string to different datatype base on input type.
// currectly support only int, uint, float, bool, and array.
// for int, uint and float: this function will always return 64 bits.
func ConvertStringTo(s string, t interface{}) interface{} {
	switch t.(type) {
	case int, int8, int16, int32, int64:
		i, e := strconv.ParseInt(s, 10, 64)
		if e == nil {
			return i
		}
	case uint, uint8, uint16, uint32, uint64:
		i, e := strconv.ParseUint(s, 10, 64)
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
		a := ForceArray(s)
		return a
	}

	return s
}

// convert string to different datatype.
// currectly support only int64, float64, and boolean
// if cannot pass it will return input string
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
