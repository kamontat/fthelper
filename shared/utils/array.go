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

	return make([]interface{}, 0), false
}
