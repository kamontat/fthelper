package xtests

type TestCase struct {
	Name     string
	Actual   interface{}
	Expected interface{}
}

func NewCase(name string, act, exp interface{}) TestCase {
	return TestCase{
		Actual:   act,
		Expected: exp,
	}
}
