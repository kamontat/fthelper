package datatype

import "fmt"

func ToString(i interface{}) (string, bool) {
	s, ok := i.(string)
	return s, ok
}

func ForceString(i interface{}) string {
	s, _ := i.(string)
	return fmt.Sprintf("%v", s)
}
