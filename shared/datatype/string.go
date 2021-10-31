package datatype

import "fmt"

// ToString will try to convert interface{} to string
func ToString(i interface{}) (string, bool) {
	s, ok := i.(string)
	return s, ok
}

// ForceString will force to convert interface{} to string
func ForceString(i interface{}) string {
	if s, ok := i.(string); ok {
		return s
	}

	return fmt.Sprintf("%v", i)
}
