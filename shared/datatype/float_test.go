package datatype_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestToFloat(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("convert normal float64 value").
		WithExpected(float64(5)).
		WithActualAndBool(datatype.ToFloat(float64(5))).
		MustEqual()
	assertion.NewName("convert normal float32 value").
		WithExpected(float64(3.44)).
		WithActualAndBool(datatype.ToFloat(float32(3.44))).
		MustEqualFloat()

	assertion.NewName("cannot convert other type (int)").
		WithActualAndBool(datatype.ToFloat(1)).
		MustError()
	assertion.NewName("cannot convert other type (bool)").
		WithActualAndBool(datatype.ToFloat(false)).
		MustError()
}

func TestForceFloat(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("convert float64").
		WithExpected(float64(5)).
		WithActualAndBool(datatype.ForceFloat(float64(5))).
		MustEqual()

	assertion.NewName("convert float string").
		WithExpected(float64(5.5)).
		WithActualAndBool(datatype.ForceFloat("5.5")).
		MustEqual()

	assertion.NewName("convert int string").
		WithExpected(float64(3)).
		WithActualAndBool(datatype.ForceFloat("3")).
		MustEqual()

	assertion.NewName("convert negative int string").
		WithExpected(float64(-55)).
		WithActualAndBool(datatype.ForceFloat("-55")).
		MustEqual()

	assertion.NewName("convert invalid string").
		WithActualAndBool(datatype.ForceFloat("-")).
		MustError()
}
