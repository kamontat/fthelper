package utils_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/utils"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestString(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("joining string").
		WithExpected("a.b.c").
		WithActual(utils.JoinString(".", "a", "b", "c")).
		MustEqual()

	assertion.NewName("joining partial empty string").
		WithExpected("a-c").
		WithActual(utils.JoinString("-", "a", "", "c")).
		MustEqual()

	assertion.NewName("joining empty string").
		WithExpected("").
		WithActual(utils.JoinString("-", "", "", "")).
		MustEqual()
}
