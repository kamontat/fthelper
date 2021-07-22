package datatype_test

import (
	"fmt"
	"testing"

	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestToString(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("convert string").
		WithExpected("str").
		WithActualAndBool(datatype.ToString("str")).
		MustEqual()

	assertion.NewName("cannot convert int").
		WithActualAndBool(datatype.ToString(1)).
		MustError()
}

type A struct {
	B bool
	F float64
}

type B struct {
	a int
	B bool
	s string
	F float64
}

func (b B) String() string {
	return fmt.Sprintf("- %d %t %s %f -", b.a, b.B, b.s, b.F)
}

func TestForceString(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("convert string").
		WithExpected("55").
		WithActual(datatype.ForceString("55")).
		MustEqual()

	assertion.NewName("convert int").
		WithExpected("5").
		WithActual(datatype.ForceString(5)).
		MustEqual()

	assertion.NewName("convert bool").
		WithExpected("true").
		WithActual(datatype.ForceString(true)).
		MustEqual()

	assertion.NewName("convert float").
		WithExpected("7.12").
		WithActual(datatype.ForceString(7.12)).
		MustEqual()

	assertion.NewName("convert empty array").
		WithExpected("[]").
		WithActual(datatype.ForceString([]string{})).
		MustEqual()

	assertion.NewName("convert array").
		WithExpected("[a b c]").
		WithActual(datatype.ForceString([]string{"a", "b", "c"})).
		MustEqual()

	assertion.NewName("convert empty map").
		WithExpected("map[]").
		WithActual(datatype.ForceString(map[string]int64{})).
		MustEqual()

	assertion.NewName("convert map").
		WithExpected("map[a:1 b:2]").
		WithActual(datatype.ForceString(map[string]int64{"a": 1, "b": 2})).
		MustEqual()

	assertion.NewName("convert struct").
		WithExpected("{true -5.5}").
		WithActual(datatype.ForceString(A{B: true, F: -5.5})).
		MustEqual()

	assertion.NewName("convert struct with String").
		WithExpected("- 10 false text 7.050000 -").
		WithActual(datatype.ForceString(B{a: 10, s: "text", B: false, F: 7.05})).
		MustEqual()
}
