package xtests

import "testing"

func New(t *testing.T) *Assertions {
	return &Assertions{
		T: t,
	}
}
