package configs_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtests"
)

func toMap(key, value string, ok bool) (maps.Mapper, bool) {
	if ok {
		return maps.New().Set(key, value), true
	}
	return maps.New(), false
}

func TestOverrideParser(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("regular string format").
		WithActualAndBool(toMap(configs.ParseOverride("test=true"))).
		WithExpected(maps.New().Set("test", "true")).
		MustDeepEqual()
	assertion.NewName("with other sign").
		WithActualAndBool(toMap(configs.ParseOverride("%$^=(!)@$"))).
		WithExpected(maps.New().Set("%$^", "(!)@$")).
		MustDeepEqual()
	assertion.NewName("too many equal sign 1").
		WithActualAndBool(toMap(configs.ParseOverride("a=b=c"))).
		MustError()
	assertion.NewName("too many equal sign 2").
		WithActualAndBool(toMap(configs.ParseOverride("a==b"))).
		MustError()
	assertion.NewName("no equal sign").
		WithActualAndBool(toMap(configs.ParseOverride("abc"))).
		MustError()
}

func TestClusterConfig(t *testing.T) {
	var assertion = xtests.New(t)

	var baseConfig = maps.New().
		Set("test", true).
		Set("status", "validated").
		Set("_.C1.test", false).
		Set("_.c1.status", "invalid").
		Set("test", true).
		Set("test", true)

	assertion.NewName("get regular config").
		WithActualAndError(configs.BuildClusterConfig("", baseConfig).Get("test")).
		WithExpected(true).
		MustEqual()
	assertion.NewName("get cluster regular config").
		WithActualAndError(configs.BuildClusterConfig("C1", baseConfig).Get("test")).
		WithExpected(false).
		MustEqual()
	assertion.NewName("get cluster regular config (case-insensitive)").
		WithActualAndError(configs.BuildClusterConfig("c1", baseConfig).Get("test")).
		WithExpected(false).
		MustEqual()
}
