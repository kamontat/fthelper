package runners_test

import (
	"errors"
	"testing"

	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestSimpleRunner(t *testing.T) {
	var assertion = xtests.New(t)

	runner := runners.NewRunner("test", runners.NoValidate, func(i *runners.SingleInfo) error {
		return nil
	})

	assertion.NewName("default status is initial").
		WithActual(runner.Information().Status()).
		WithExpected(runners.INITIAL).
		MustEqual()

	assertion.NewName("should be valid").
		WithError(runner.Validate()).
		MustNotError()

	assertion.NewName("should be runnable").
		WithError(runner.Run()).
		MustNotError()

	assertion.NewName("after run, status is success").
		WithActual(runner.Information().Status()).
		WithExpected(runners.SUCCESS).
		MustEqual()

	assertion.NewName("after run, duration is increase").
		WithActual(runner.Information().Duration()).
		WithExpected(0).
		MustGreaterThan()

	assertion.NewName("after run, no error").
		WithError(runner.Information().Error().Error()).
		MustNotError()
}

func requireInput(i *runners.SingleInfo) error {
	if i.Input() == nil {
		return errors.New("input is required")
	}
	return nil
}

func TestInputRunner(t *testing.T) {
	var assertion = xtests.New(t)
	runnerWithInput := runners.NewRunner("input", requireInput, runners.NoValidate).Input("hello world")
	runnerWithoutInput := runners.NewRunner("no input", requireInput, runners.NoValidate)

	assertion.NewName("correct initial input value").
		WithError(runnerWithInput.Run()).
		MustNotError()

	assertion.NewName("input is required error").
		WithError(runnerWithoutInput.Validate()).
		WithExpected("input is required").
		MustContainError()

	assertion.NewName("validate didn't rerun").
		WithError(runnerWithoutInput.Validate()).
		WithExpected("input is required").
		MustContainError()

	assertion.NewName("run return error same as validate").
		WithError(runnerWithoutInput.Run()).
		WithExpected("input is required").
		MustContainError()
}
