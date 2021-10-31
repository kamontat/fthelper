package datatype

import "strconv"

// ToFloat will try to convert interface{} to float
func ToFloat(i interface{}) (float64, bool) {
	f32, ok := i.(float32)
	if ok {
		return float64(f32), ok
	}

	f64, ok := i.(float64)
	if ok {
		return f64, ok
	}

	return -1, false
}

// ForceFloat will force to convert interface{} to float
func ForceFloat(i interface{}) (float64, bool) {
	var f, ok = ToFloat(i)
	if ok {
		return f, ok
	}

	var str = ForceString(i)
	var f64, err = strconv.ParseFloat(str, 64)
	if err == nil {
		return f64, true
	}

	return -1, false
}
