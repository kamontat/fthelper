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

	assertion.NewName("trim empty string").
		WithExpected("").
		WithActual(utils.TrimString("", 8)).
		MustEqual()
	assertion.NewName("trim short string").
		WithExpected("short").
		WithActual(utils.TrimString("short", 5)).
		MustEqual()
	assertion.NewName("trim long string when not exceed dot syntax").
		WithExpected("interes...").
		WithActual(utils.TrimString("interesting", 10)).
		MustEqual()
	assertion.NewName("trim long string when exceed dot syntax").
		WithExpected("this i...").
		WithActual(utils.TrimString("this is very long string", 9)).
		MustEqual()

	assertion.NewName("not mask at all").
		WithExpected("secure string").
		WithActual(utils.MaskString("secure string", utils.NONE)).
		MustEqual()
	assertion.NewName("mask low").
		WithExpected("secure st****").
		WithActual(utils.MaskString("secure string", utils.LOW)).
		MustEqual()
	assertion.NewName("mask medium").
		WithExpected("sec*********g").
		WithActual(utils.MaskString("secure string", utils.MEDIUM)).
		MustEqual()
	assertion.NewName("mask high").
		WithExpected("sec****g").
		WithActual(utils.MaskString("secure string", utils.HIGH)).
		MustEqual()
}
