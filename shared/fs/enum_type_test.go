package fs_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestEnumType(t *testing.T) {
	var assertion = xtests.New(t)
	assertion.NewName("empty type").
		WithActualAndBool(fs.ToType("")).
		MustError()
	assertion.NewName("auto type").
		WithExpected(fs.AUTO).
		WithActualAndBool(fs.ToType("auto")).
		MustEqual()
	assertion.NewName("auto type (ignore case)").
		WithExpected(fs.AUTO).
		WithActualAndBool(fs.ToType("AutO")).
		MustEqual()
	assertion.NewName("file type").
		WithExpected(fs.FILE).
		WithActualAndBool(fs.ToType("file")).
		MustEqual()
	assertion.NewName("file type (short)").
		WithExpected(fs.FILE).
		WithActualAndBool(fs.ToType("F")).
		MustEqual()
	assertion.NewName("file type (ignore case)").
		WithExpected(fs.FILE).
		WithActualAndBool(fs.ToType("fIlE")).
		MustEqual()
	assertion.NewName("directory type").
		WithExpected(fs.DIRECTORY).
		WithActualAndBool(fs.ToType("directory")).
		MustEqual()
	assertion.NewName("directory type (short)").
		WithExpected(fs.DIRECTORY).
		WithActualAndBool(fs.ToType("dIR")).
		MustEqual()
	assertion.NewName("directory type (ignore case)").
		WithExpected(fs.DIRECTORY).
		WithActualAndBool(fs.ToType("diReCToRY")).
		MustEqual()
}
