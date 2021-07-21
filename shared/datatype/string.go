package datatype

import "fmt"

func ToString(i interface{}) (string, bool) {
	s, ok := i.(string)
	return s, ok
}

func ForceString(i interface{}) string {
	if s, ok := i.(string); ok {
		return s
	}

	return fmt.Sprintf("%v", i)
}
