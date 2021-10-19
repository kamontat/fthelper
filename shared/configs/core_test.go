package configs_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtests"
)

// TODO: This test didn't works as I expected, Actually this throw error out because config builder will try to query data from config files/directories as well
func TestCore(t *testing.T) {
	var assertion = xtests.New(t)

	var builder = configs.New("config", maps.New().Set("test", true))
	assertion.NewName("direct configuration").
		WithActualAndError(builder.Build([]string{})).
		WithExpected(maps.New().Set("test", true)).
		MustDeepEqual()
}
