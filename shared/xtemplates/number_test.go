package xtemplates_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtemplates"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestRatio(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("ratio - 1%").
		WithExpected(`0.01`).
		WithActualAndError(xtemplates.Text(`{{ ratio "1%" }}`, maps.New())).
		MustEqual()
	assertion.NewName("ratio - 101%").
		WithExpected(`1.01`).
		WithActualAndError(xtemplates.Text(`{{ ratio "101%" }}`, maps.New())).
		MustEqual()
	assertion.NewName("ratio - 100").
		WithExpected(`1`).
		WithActualAndError(xtemplates.Text(`{{ ratio "100" }}`, maps.New())).
		MustEqual()
	assertion.NewName("wrong ratio").
		WithActualAndError(xtemplates.Text(`{{ ratio "100a" }}`, maps.New())).
		MustError()
}
