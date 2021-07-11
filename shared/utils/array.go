package utils

// CloneArray will create copy of input array
func CloneArray(a []interface{}, extra ...interface{}) []interface{} {
	var base = make([]interface{}, 0)

	base = append(base, a...)
	base = append(base, extra...)

	return base
}

func CloneStringArray(a []string, extra ...string) []string {
	var base = make([]string, 0)

	base = append(base, a...)
	base = append(base, extra...)

	return base
}
