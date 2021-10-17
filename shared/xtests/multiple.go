package xtests

import "testing"

type Assertions struct {
	T *testing.T
}

func (a *Assertions) New() *Assertion {
	return &Assertion{
		enabled:  true,
		name:     "",
		desc:     "",
		actual:   nil,
		expected: nil,
		err:      nil,
		T:        a.T,
	}
}

func (a *Assertions) NewName(name string) *Assertion {
	return &Assertion{
		enabled:  true,
		name:     name,
		desc:     "",
		actual:   nil,
		expected: nil,
		err:      nil,
		T:        a.T,
	}
}
