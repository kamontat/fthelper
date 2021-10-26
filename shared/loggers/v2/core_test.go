package loggers_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/loggers/v2"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestGlobalLevel(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("Default global level is INFO").
		WithActual(loggers.GetLevel()).
		WithExpected(loggers.INFO).
		MustEqual()

	loggers.SetLevel(4)
	assertion.NewName("Able to setup global level via SetLevel").
		WithActual(loggers.GetLevel()).
		WithExpected(loggers.DEBUG).
		MustEqual()

	loggers.SetLevel(-72)
	assertion.NewName("SetLevel should round if user enter wrong level").
		WithActual(loggers.GetLevel()).
		WithExpected(loggers.SILENT).
		MustEqual()
}
