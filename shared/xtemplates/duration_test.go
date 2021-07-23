package xtemplates_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtemplates"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestDurationFunction(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("normal duration - hour").
		WithExpected("0.25").
		WithActualAndError(xtemplates.Text("{{ toDuration `15m` `h` }}", maps.New())).
		MustEqual()

	assertion.NewName("normal duration - minute").
		WithExpected("5").
		WithActualAndError(xtemplates.Text("{{ toDuration `5m` `m` }}", maps.New())).
		MustEqual()

	assertion.NewName("normal duration - second").
		WithExpected("300").
		WithActualAndError(xtemplates.Text("{{ toDuration `5m` `s` }}", maps.New())).
		MustEqual()

	assertion.NewName("normal duration - millisecond").
		WithExpected("300000").
		WithActualAndError(xtemplates.Text("{{ toDuration `5m` `ms` }}", maps.New())).
		MustEqual()

	assertion.NewName("day duration - second").
		WithExpected("86400").
		WithActualAndError(xtemplates.Text("{{ dayToDuration 1 `s` }}", maps.New())).
		MustEqual()

	assertion.NewName("empty input").
		WithExpected("time: invalid duration \"\"").
		WithActualAndError(xtemplates.Text("{{ toDuration `` `s` }}", maps.New())).
		Must(xtests.MUST_ERROR, xtests.MUST_CONTAINS_ERROR)

	assertion.NewName("invalid unit").
		WithExpected("cannot convert to unit 'a'").
		WithActualAndError(xtemplates.Text("{{ toDuration `5m` `a` }}", maps.New())).
		Must(xtests.MUST_ERROR, xtests.MUST_CONTAINS_ERROR)
}
