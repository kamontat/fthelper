package runners_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/utils"
)

func mockRunner() runners.Runner {
	return runners.NewRunner(utils.RandString(8), runners.NoValidate, func(i *runners.SingleInfo) error {
		return nil
	})
}

func TestCreateSummary(t *testing.T) {
	t.Run("create simple summary from collection", func(t *testing.T) {
		c := runners.NewCollection("test", mockRunner(), mockRunner(), mockRunner())
		summary := runners.ColSummary(c)

		summary.Log(loggers.Get("test"))
	})
}
