package fs_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestEnumMode(t *testing.T) {
	var assertion = xtests.New(t)
	assertion.NewName("empty mode").
		WithActualAndBool(fs.ToMode("")).
		MustError()
	assertion.NewName("single mode").
		WithExpected(fs.SINGLE).
		WithActualAndBool(fs.ToMode("single")).
		MustEqual()
	assertion.NewName("single mode (short)").
		WithExpected(fs.SINGLE).
		WithActualAndBool(fs.ToMode("S")).
		MustEqual()
	assertion.NewName("single mode (ignore case)").
		WithExpected(fs.SINGLE).
		WithActualAndBool(fs.ToMode("sIngle")).
		MustEqual()
	assertion.NewName("multiple mode").
		WithExpected(fs.MULTIPLE).
		WithActualAndBool(fs.ToMode("multiple")).
		MustEqual()
	assertion.NewName("multiple mode (short)").
		WithExpected(fs.MULTIPLE).
		WithActualAndBool(fs.ToMode("m")).
		MustEqual()
	assertion.NewName("multiple mode (ignore case)").
		WithExpected(fs.MULTIPLE).
		WithActualAndBool(fs.ToMode("MULTIPLE")).
		MustEqual()
}
