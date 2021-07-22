package datatype_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestToInt(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("convert int64").
		WithExpected(int64(99)).
		WithActualAndBool(datatype.ToInt(int64(99))).
		MustEqual()
	assertion.NewName("convert int").
		WithExpected(int64(5)).
		WithActualAndBool(datatype.ToInt(5)).
		MustEqual()
	assertion.NewName("convert int8").
		WithExpected(int64(1)).
		WithActualAndBool(datatype.ToInt(int8(1))).
		MustEqual()
	assertion.NewName("convert int16").
		WithExpected(int64(2)).
		WithActualAndBool(datatype.ToInt(int16(2))).
		MustEqual()
	assertion.NewName("convert int32").
		WithExpected(int64(12)).
		WithActualAndBool(datatype.ToInt(int32(12))).
		MustEqual()

	assertion.NewName("cannot convert other type (int)").
		WithActualAndBool(datatype.ToInt(1.5)).
		MustError()
	assertion.NewName("cannot convert other type (bool)").
		WithActualAndBool(datatype.ToInt(false)).
		MustError()
}

func TestForceInt(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("convert int64").
		WithExpected(int64(99)).
		WithActualAndBool(datatype.ForceInt(int64(99))).
		MustEqual()

	assertion.NewName("convert int string").
		WithExpected(int64(5)).
		WithActualAndBool(datatype.ForceInt("5")).
		MustEqual()

	assertion.NewName("cannot convert other type (int)").
		WithActualAndBool(datatype.ForceInt(1.5)).
		MustError()

	assertion.NewName("cannot convert other type (bool)").
		WithActualAndBool(datatype.ForceInt(false)).
		MustError()
}
