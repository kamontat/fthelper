package datatype_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestToBool(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("convert normal boolean value").
		WithExpected(true).
		WithActualAndBool(datatype.ToBool(true)).
		MustEqual()

	assertion.NewName("cannot convert other type (int)").
		WithActualAndBool(datatype.ToBool(1)).
		MustError()

	assertion.NewName("cannot convert other type (string)").
		WithActualAndBool(datatype.ToBool("true")).
		MustError()
}

func TestForceBool(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("convert normal boolean value").
		WithExpected(true).
		WithActualAndBool(datatype.ForceBool(true)).
		MustEqual()

	assertion.NewName("convert number one (1)").
		WithExpected(true).
		WithActualAndBool(datatype.ForceBool(1)).
		MustEqual()

	assertion.NewName("convert string boolean").
		WithExpected(true).
		WithActualAndBool(datatype.ForceBool("true")).
		MustEqual()

	assertion.NewName("cannot convert minus number").
		WithActualAndBool(datatype.ForceBool(-1)).
		MustError()
}
