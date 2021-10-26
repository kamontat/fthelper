package loggers_test

import (
	"testing"

	loggers "github.com/kamontat/fthelper/shared/loggers/v2"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestToLevel(t *testing.T) {
	var assertion = xtests.New(t)
	for _, tc := range []xtests.TestCase{
		xtests.NewCase("silent level", float64(0), loggers.SILENT),
		xtests.NewCase("error level", float64(1), loggers.ERROR),
		xtests.NewCase("warn level", float64(2), loggers.WARN),
		xtests.NewCase("info level", float64(3), loggers.INFO),
		xtests.NewCase("debug level", float64(4), loggers.DEBUG),
		xtests.NewCase("negative level", float64(-20), loggers.SILENT),
		xtests.NewCase("exceeded level", float64(99), loggers.DEBUG),
	} {
		assertion.NewName(tc.Name).
			WithActual(loggers.ToLevel(tc.Actual.(float64))).
			WithExpected(tc.Expected).
			MustEqual()
	}
}
