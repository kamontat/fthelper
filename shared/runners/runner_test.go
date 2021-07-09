package runners_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/runners"
)

func TestCreateRunner(t *testing.T) {
	t.Run("create simple runner", func(t *testing.T) {
		runner := runners.NewRunner("test", runners.NoValidate, func(i *runners.SingleInfo) error {
			return nil
		})

		if err := runner.Validate(); err != nil {
			t.Errorf("validate should return nil error")
		}

		if err := runner.Run(); err != nil {
			t.Errorf("run should return nil error")
		}

		if runner.Information().Status() != runners.SUCCESS {
			t.Errorf("information must updated to success status")
		}
		if runner.Information().Duration() <= 0 {
			t.Errorf("information must update duration")
		}
		if runner.Information().Error().HasError() {
			t.Errorf("information must not update error")
		}
	})
}
