package loggers_test

import (
	"bytes"
	"testing"

	loggers "github.com/kamontat/fthelper/shared/loggers/v2"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestNewPrinter(t *testing.T) {
	var assertion = xtests.New(t)

	var buffer1 *bytes.Buffer = &bytes.Buffer{}
	var printer = loggers.NewPrinter(buffer1)

	printer.Print("test message")
	assertion.NewName("Print output correct string to builtin writer").
		WithActual(buffer1.String()).
		WithExpected("test message\n").
		MustEqual()

	var buffer2 *bytes.Buffer = &bytes.Buffer{}
	printer.Write(buffer2, "second message")
	assertion.NewName("Write output correct string to custom writer").
		WithActual(buffer2.String()).
		WithExpected("second message\n").
		MustEqual()
}
