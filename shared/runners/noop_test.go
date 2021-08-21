package runners_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestNoopFunction(t *testing.T) {
	var assertion = xtests.New(t)
	assertion.New().
		WithActual(runners.NoValidate(&runners.SingleInfo{})).
		MustBeNil()
}
