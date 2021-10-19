package configs_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestParseConfigFromEnv(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("normal environment").
		WithActualAndError(configs.ParseConfigFromEnv([]string{"FTH_FREQTRADE_STATUS=running"})).
		WithExpected(maps.New().Set("freqtrade-status", "running")).
		MustDeepEqual()
	assertion.NewName("unknown environment").
		WithActualAndError(configs.ParseConfigFromEnv([]string{"FTHSTATUS=running"})).
		WithExpected(maps.New()).
		MustDeepEqual()
}
