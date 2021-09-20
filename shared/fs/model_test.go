package fs_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestConstant(t *testing.T) {
	var assertion = xtests.New(t)
	assertion.NewName("verify path separator '/'").
		WithActual(fs.Separator).
		WithExpected("/").
		MustEqual()
}

func TestFile(t *testing.T) {
	var assertion = xtests.New(t)

	var file, err = fs.NewFile([]string{"tmp", "test", "readme.txt"})
	assertion.NewName("create without error").
		WithActualAndError(file, err).
		MustNotError()
	assertion.NewName("get basename correctly").
		WithActual(file.Basename()).
		WithExpected("readme.txt").
		MustEqual()
	assertion.NewName("get name correctly").
		WithActual(file.Name()).
		WithExpected("readme").
		MustEqual()
	assertion.NewName("get directory path correctly").
		WithActual(file.Dirpath()).
		WithExpected("/tmp/test").
		MustEqual()
	assertion.NewName("get directory name correctly").
		WithActual(file.Dirname()).
		WithExpected("test").
		MustEqual()
	assertion.NewName("get absolute path correctly").
		WithActual(file.Abs()).
		WithExpected("/tmp/test/readme.txt").
		MustEqual()
}
