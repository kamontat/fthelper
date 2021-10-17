package errors_test

import (
	"fmt"
	"testing"

	"github.com/kamontat/fthelper/shared/errors"
	"github.com/kamontat/fthelper/shared/utils"
	"github.com/kamontat/fthelper/shared/xtests"
)

func Mock() error {
	return String(utils.RandString(10))
}

func String(s string) error {
	return fmt.Errorf(s)
}

func TestErrorHandler(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("empty handler").
		WithActual(errors.New()).
		MustNotBeNil()

	assertion.NewName("error length").
		WithExpected(1).
		WithActual(errors.New().And(Mock()).And(nil).Length()).
		MustEqual()

	assertion.NewName("and error").
		WithExpected(true).
		WithActual(errors.New().And(Mock()).And(nil).HasError()).
		MustEqual()

	var data, err = errors.New().AndD("string", Mock())
	assertion.NewName("andD return error").
		WithExpected(true).
		WithActual(err.HasError()).
		MustEqual()
	assertion.NewName("andD return data").
		WithExpected("string").
		WithActual(data).
		MustEqual()

	assertion.NewName("merging").
		WithExpected(2).
		WithActual(errors.New().And(Mock()).Merge(errors.New().And(Mock())).Length()).
		MustEqual()

	var first = Mock()
	assertion.NewName("get first").
		WithExpected(first).
		WithActual(errors.New().And(first).Merge(errors.New().And(Mock())).First()).
		MustEqual()

	assertion.NewName("get first nil").
		WithActual(errors.New().First()).
		MustBeNil()

	assertion.NewName("string - with error").
		WithExpected(`found '2' errors (total=2)
- a
- b
`).
		WithActual(errors.New().And(String("a")).And(String("b")).String()).
		MustEqual()

	assertion.NewName("string - not error").
		WithExpected(`not found any errors`).
		WithActual(errors.New().String()).
		MustEqual()

	assertion.NewName("get error - when no error").
		WithActual(errors.New().Error()).
		MustBeNil()

	assertion.NewName("get error - when has error").
		WithExpected("this is error message").
		WithError(errors.New().And(fmt.Errorf("this is error message")).Error()).
		MustEqualError()
}
