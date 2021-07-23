package xtests_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/xtests"
)

func TestXTest(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.New().WithActual(assertion).MustNotBeNil()

	assertion.NewName(xtests.MUST_EQUAL).
		WithExpected(true).
		WithActual(true).
		MustEqual()

	assertion.NewName(xtests.MUST_EQUAL_FLOAT).
		WithExpected(0.1234568).
		WithActual(0.1234567).
		MustEqualFloat()

	assertion.NewName(xtests.MUST_DEEP_EQUAL).
		WithExpected([]int{1, 2, 3}).
		WithActualAndBool([]int{1, 2, 3}, true).
		MustDeepEqual()

	assertion.NewName(xtests.MUST_BE_NIL).
		WithActualAndBool(nil, false).
		MustBeNil()

	assertion.NewName(xtests.MUST_BE_NIL+xtests.MUST_ERROR).
		WithActualAndBool(nil, false).
		Must(xtests.MUST_BE_NIL, xtests.MUST_ERROR)
}
