package runners_test

import (
	"testing"
	"time"

	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestSingleInfo(t *testing.T) {
	var assertion = xtests.New(t)

	var info = runners.NewSingleInfo("test")
	assertion.NewName("create new single info").
		WithActual(info).
		MustNotBeNil()

	assertion.NewName("get correct name").
		WithExpected("test").
		WithActual(info.Name()).
		MustEqual()

	assertion.NewName("get default duration").
		WithExpected(time.Duration(0)).
		WithActual(info.Duration()).
		MustEqual()

	assertion.NewName("get no error").
		WithExpected(false).
		WithActual(info.Error().HasError()).
		MustEqual()

	assertion.NewName("get nil input").
		WithActual(info.Input()).
		MustBeNil()

	assertion.NewName("get input as 120").
		WithExpected(120).
		WithActual(info.In(120).Input()).
		MustEqual()

	assertion.NewName("get output as nil").
		WithActual(info.Output()).
		MustBeNil()

	assertion.NewName("get output as false").
		WithExpected(false).
		WithActual(info.Out(false).Output()).
		MustEqual()

	assertion.NewName("default status is initial").
		WithExpected(runners.INITIAL).
		WithActual(info.Status()).
		MustEqual()

	assertion.NewName("total count always be 1").
		WithExpected(1).
		WithActual(info.TotalCount()).
		MustEqual()
}
