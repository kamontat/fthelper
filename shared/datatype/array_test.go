package datatype_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestToArray(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("create from interface array").
		WithExpected([]interface{}{"raw", "value"}).
		WithActualAndBool(datatype.ToArray([]interface{}{"raw", "value"})).
		MustDeepEqual()

	assertion.NewName("create from string array").
		WithExpected([]interface{}{"1", "2"}).
		WithActualAndBool(datatype.ToArray([]string{"1", "2"})).
		MustDeepEqual()

	assertion.NewName("create from int array").
		WithExpected([]interface{}{5, 6}).
		WithActualAndBool(datatype.ToArray([]int{5, 6})).
		MustDeepEqual()

	assertion.NewName("create from string").
		WithExpected([]interface{}{"1", "2", "3"}).
		WithActualAndBool(datatype.ToArray("a:1,2,3")).
		MustDeepEqual()

	assertion.NewName("create from alternative string").
		WithExpected([]interface{}{"true", "false"}).
		WithActualAndBool(datatype.ToArray("array:true,false")).
		MustDeepEqual()

	assertion.NewName("input invalid string format").
		WithActualAndBool(datatype.ToArray("true,false")).
		MustError()

	assertion.NewName("input invalid value").
		WithActualAndBool(datatype.ToArray(500)).
		MustError()
}

func TestForceArray(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("create from interface array").
		WithExpected([]interface{}{"raw", "value"}).
		WithActual(datatype.ForceArray([]interface{}{"raw", "value"})).
		MustDeepEqual()

	assertion.NewName("without indicator but still pass").
		WithExpected([]interface{}{"one", "two"}).
		WithActual(datatype.ForceArray("one,two")).
		MustDeepEqual()

	assertion.NewName("string will be one element array").
		WithExpected([]interface{}{"string"}).
		WithActual(datatype.ForceArray("string")).
		MustDeepEqual()

	assertion.NewName("int will be one element array").
		WithExpected([]interface{}{500}).
		WithActual(datatype.ForceArray(500)).
		MustDeepEqual()
}
