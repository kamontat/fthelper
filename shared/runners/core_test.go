package runners_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/utils"
)

func TestRunner(t *testing.T) {
	var rs = runners.New()
	for range make([]int, 25) {
		rs.Add(runners.NewRunner(utils.RandString(8), runners.NoValidator, runners.NoExecutor))
	}

	// rs.Add(runners.NewRunner("first runner", func(context *runners.Context) error {
	// 	return errors.New("test")
	// }, func(context *runners.Context) error {
	// 	return nil
	// }), runners.NewRunner("fix me please", func(context *runners.Context) error {
	// 	return nil
	// }, func(context *runners.Context) error {
	// 	return errors.New("test")
	// }))
	// .AddGroup("morning group", runners.NewRunner("another please", func(context *runners.Context) error {
	// 	return nil
	// }, runners.NoExecutor), runners.NewRunner("computer engineer", runners.NoValidator, runners.NoExecutor))

	var summary = rs.Run()
	summary.Log(loggers.Get())
}
