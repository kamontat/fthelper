package xtemplates_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtemplates"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestJoin(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("joining string").
		WithExpected("a-b-c").
		WithActualAndError(xtemplates.Text(`{{ join "a" "b" "c" }}`, maps.New())).
		MustEqual()

	assertion.NewName("joining partial empty string").
		WithExpected("a-c").
		WithActualAndError(xtemplates.Text(`{{ join "a" "" "c" }}`, maps.New())).
		MustEqual()

	assertion.NewName("joining empty string").
		WithExpected("").
		WithActualAndError(xtemplates.Text(`{{ join "" "" "" }}`, maps.New())).
		MustEqual()

	assertion.NewName("joining empty array").
		WithExpected("").
		WithActualAndError(xtemplates.Text(`{{ join }}`, maps.New())).
		MustEqual()
}
