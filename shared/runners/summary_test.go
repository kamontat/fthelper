package runners_test

import (
	"bytes"
	"testing"

	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/utils"
	"github.com/kamontat/fthelper/shared/xtests"
)

func mockRunner() runners.Runner {
	return runners.NewRunner(utils.RandString(8), runners.NoValidate, func(i *runners.SingleInfo) error {
		return nil
	})
}

func mockLogger() (*loggers.Logger, *bytes.Buffer) {
	var buffer = new(bytes.Buffer)
	var logger = loggers.Get("test")

	logger.SetWriter(buffer)
	return logger, buffer
}

func TestCreateSummary(t *testing.T) {
	var assertion = xtests.New(t)
	t.Run("create simple summary from collection", func(t *testing.T) {
		var logger, writer = mockLogger()
		c := runners.NewCollection("test", mockRunner(), mockRunner(), mockRunner())
		summary := runners.ColSummary(c)
		summary.Log(logger)

		assertion.New().
			WithActual(writer.String()).
			WithExpected("SUMMARY 3 TASKS: success").
			MustContain()
	})

	t.Run("create summary from runner", func(t *testing.T) {
		var logger, writer = mockLogger()
		summary := runners.RunSummary(mockRunner())
		summary.Log(logger)

		assertion.New().
			WithActual(writer.String()).
			WithExpected("success ( 01/01 - 100.00%)").
			MustContain()
	})
}
