package xtemplates_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtemplates"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestDurationFunction(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("normal duration").
		WithExpected("300").
		WithActualAndError(xtemplates.Text("{{ toDuration `5m` `s` }}", maps.New())).
		Must(xtests.MUST_NOT_ERROR, xtests.MUST_EQUAL)

	assertion.NewName("empty input").
		WithExpected("time: invalid duration \"\"").
		WithActualAndError(xtemplates.Text("{{ toDuration `` `s` }}", maps.New())).
		Must(xtests.MUST_ERROR, xtests.MUST_CONTAINS_ERROR)

	assertion.NewName("invalid unit").
		WithExpected("cannot convert to unit 'a'").
		WithActualAndError(xtemplates.Text("{{ toDuration `5m` `a` }}", maps.New())).
		Must(xtests.MUST_ERROR, xtests.MUST_CONTAINS_ERROR)
}
