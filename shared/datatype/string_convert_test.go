package datatype_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestConvertString(t *testing.T) {
	var assertion = xtests.New(t)
	assertion.NewName("convert string to string").
		WithExpected("str").
		WithActual(datatype.ConvertString("str")).
		MustEqual()

	assertion.NewName("convert string to int64").
		WithExpected(int64(55)).
		WithActual(datatype.ConvertString("55")).
		MustEqual()

	assertion.NewName("convert string to float64").
		WithExpected(float64(5.5)).
		WithActual(datatype.ConvertString("5.5")).
		MustEqual()

	assertion.NewName("convert string to bool").
		WithExpected(true).
		WithActual(datatype.ConvertString("true")).
		MustEqual()
}

func TestConvertStringTo(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("convert string to string").
		WithExpected("str").
		WithActual(datatype.ConvertStringTo("str", "")).
		MustEqual()

	assertion.NewName("convert string to int64").
		WithExpected(int64(55)).
		WithActual(datatype.ConvertStringTo("55", 0)).
		MustEqual()

	assertion.NewName("convert string to uint64").
		WithExpected(uint64(55)).
		WithActual(datatype.ConvertStringTo("55", uint(0))).
		MustEqual()

	assertion.NewName("convert string to float64").
		WithExpected(float64(5.5)).
		WithActual(datatype.ConvertStringTo("5.5", float32(0))).
		MustEqual()

	assertion.NewName("convert string to bool").
		WithExpected(true).
		WithActual(datatype.ConvertStringTo("true", false)).
		MustEqual()

	assertion.NewName("convert to input value (string)").
		WithExpected("true").
		WithActual(datatype.ConvertStringTo("true", "string")).
		MustEqual()

	assertion.NewName("convert to input value (array)").
		WithExpected([]interface{}{"a", "b", "c"}).
		WithActual(datatype.ConvertStringTo("a,b,c", make([]string, 0))).
		MustDeepEqual()
}
