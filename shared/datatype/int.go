package datatype

import "strconv"

func ToInt(i interface{}) (int64, bool) {
	i32, ok := i.(int32)
	if ok {
		return int64(i32), ok
	}

	i64, ok := i.(int64)
	if ok {
		return i64, ok
	}

	return -1, false
}

func ForceInt(inf interface{}) (int64, bool) {
	var i, ok = ToInt(inf)
	if ok {
		return i, ok
	}

	var str = ForceString(inf)
	var i64, err = strconv.ParseInt(str, 10, 64)
	if err == nil {
		return i64, true
	}

	return -1, false
}
