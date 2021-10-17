package xtests

type TestCase struct {
	Name     string
	Actual   interface{}
	Expected interface{}
	Checker  []MustChecker
}

func NewCase(name string, act, exp interface{}, checker ...MustChecker) TestCase {
	return TestCase{
		Actual:   act,
		Expected: exp,
		Checker:  checker,
	}
}
