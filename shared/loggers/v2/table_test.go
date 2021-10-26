package loggers_test

import (
	"bytes"
	"testing"

	"github.com/kamontat/fthelper/shared/loggers/v2"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestTableError(t *testing.T) {
	var assertion = xtests.New(t)

	var buffer1 = &bytes.Buffer{}
	var printer = loggers.NewPrinter(buffer1)
	var table = loggers.NewTable(printer)

	assertion.NewName("error from Row()").
		WithError(table.Row("hello", "world")).
		WithExpected("you never initial table").
		MustContainError()
	assertion.NewName("empty output").
		WithActual(buffer1.String()).
		WithExpected("").
		MustEqual()
	assertion.NewName("silent end").
		WithError(table.End()).
		MustNotError()
}

func TestNewTable(t *testing.T) {
	var assertion = xtests.New(t)

	var buffer1 = &bytes.Buffer{}
	var printer = loggers.NewPrinter(buffer1)
	var table = loggers.NewTable(printer)
	table.SetSize(2).Init()

	assertion.NewName("not error").
		WithError(table.Row("hello", "world")).
		MustNotError()
	assertion.NewName("able to call End()").
		WithError(table.End()).
		MustNotError()
	assertion.NewName("create new row").
		WithActual(buffer1.String()).
		WithExpected(`hello  world
`).
		MustEqual()
}
