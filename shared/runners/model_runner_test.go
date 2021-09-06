package runners_test

import (
	"errors"
	"testing"

	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestBasicRunner(t *testing.T) {
	var assertion = xtests.New(t)

	var runner = runners.NewRunner("test", runners.NoValidator, runners.NoExecutor)
	assertion.NewName("get runner name").
		WithExpected("test").
		WithActual(runner.Name).
		MustEqual()
	assertion.NewName("enabled by default").
		WithExpected(false).
		WithActual(runner.Context.Disabled).
		MustEqual()
	assertion.NewName("did not have input value").
		WithExpected(nil).
		WithActual(runner.Context.Input()).
		MustEqual()

	assertion.NewName("get result information").
		WithExpected("^test: success \\(([0-9]+[nµ]s)\\)$").
		WithActual(runner.Run().String()).
		MustEqualRegex()
}

func TestInvalidRunner(t *testing.T) {
	var assertion = xtests.New(t)

	var runner = runners.NewRunner("name", func(context *runners.Context) error {
		return errors.New("error message")
	}, runners.NoExecutor)
	assertion.NewName("get result information").
		WithExpected("^name: invalid \\(([0-9]+[nµ]s)\\)$").
		WithActual(runner.Run().String()).
		MustEqualRegex()
}

func TestErrorRunner(t *testing.T) {
	var assertion = xtests.New(t)

	var runner = runners.NewRunner("information", runners.NoValidator, func(context *runners.Context) error {
		return errors.New("error message")
	})
	assertion.NewName("get result information").
		WithExpected("^information: error \\(([0-9]+[nµ]s)\\)$").
		WithActual(runner.Run().String()).
		MustEqualRegex()
}

func TestDisableRunner(t *testing.T) {
	var assertion = xtests.New(t)

	// You cannot call disable from executor since it already execute
	// You still able to disable in validator,
	// however disabled before .Run() command will result as skip both validator and executor

	var counter = 0
	var before = runners.NewRunner("message", func(context *runners.Context) error {
		counter++
		return nil
	}, func(context *runners.Context) error {
		counter++
		return nil
	})
	assertion.NewName("get result information").
		WithExpected("^message: disabled \\(([0-9]+[nµ]s)\\)$").
		WithActual(before.Disable(true).Run().String()).
		MustEqualRegex()
	assertion.NewName("counter is not increase").
		WithExpected(0).
		WithActual(counter).
		MustEqual()

	counter = 0
	var inValidate = runners.NewRunner("validate", func(context *runners.Context) error {
		context.Disabled = true
		counter++
		return nil
	}, func(context *runners.Context) error {
		counter++
		return nil
	})
	assertion.NewName("get result information").
		WithExpected("^validate: disabled \\(([0-9]+[nµ]s)\\)$").
		WithActual(inValidate.Run().String()).
		MustEqualRegex()
	assertion.NewName("counter increase by one").
		WithExpected(1).
		WithActual(counter).
		MustEqual()
}
