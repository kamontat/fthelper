package datatype

import "strconv"

// ToBool will try to convert interface{} to boolean
func ToBool(i interface{}) (bool, bool) {
	b, ok := i.(bool)
	if ok {
		return b, ok
	}

	return false, false
}

// ForceBool will force to convert interface{} to boolean
func ForceBool(inf interface{}) (bool, bool) {
	var b, ok = ToBool(inf)
	if ok {
		return b, ok
	}

	var str = ForceString(inf)
	var sb, err = strconv.ParseBool(str)
	if err == nil {
		return sb, true
	}

	return false, false
}
