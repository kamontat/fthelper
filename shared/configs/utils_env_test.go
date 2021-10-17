package configs_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestIsEnvKey(t *testing.T) {
	var assertion = xtests.New(t)

	for _, tc := range []xtests.TestCase{
		xtests.NewCase("normal case", "FTH_TEST", true),
		xtests.NewCase("with symbol", "FTH_ABC$%", true),
		xtests.NewCase("missing prefix", "TEST", false),
		xtests.NewCase("prefix not on prefix position", "TEST_FTH_TEST", false),
		xtests.NewCase("wrong prefix", "FTB_TEST", false),
		xtests.NewCase("prefix without separator", "FTHTEST", false),
	} {
		assertion.NewName(tc.Name).
			WithActual(configs.IsEnvKey(tc.Actual.(string))).
			WithExpected(tc.Expected).
			MustEqual()
	}
}

func TestIsCEnvKey(t *testing.T) {
	var assertion = xtests.New(t)

	for _, tc := range []xtests.TestCase{
		xtests.NewCase("normal case", "FTC_TEST", true),
		xtests.NewCase("with symbol", "FTC_ABC$%", true),
		xtests.NewCase("missing prefix", "TEST", false),
		xtests.NewCase("prefix not on prefix position", "TEST_FTC_TEST", false),
		xtests.NewCase("wrong prefix", "FTB_TEST", false),
		xtests.NewCase("prefix without separator", "FTCTEST", false),
	} {
		assertion.NewName(tc.Name).
			WithActual(configs.IsCEnvKey(tc.Actual.(string))).
			WithExpected(tc.Expected).
			MustEqual()
	}
}

func TestEnvToKey(t *testing.T) {
	var assertion = xtests.New(t)
	for _, tc := range []xtests.TestCase{
		xtests.NewCase("not environment format", "FTP_TEST", "", xtests.MUST_ERROR),
		xtests.NewCase("normal case (cenv)", "FTC_TEST", "_.test", xtests.MUST_EQUAL),
		xtests.NewCase("normal case (env)", "FTH_TEST", "test", xtests.MUST_EQUAL),
		xtests.NewCase("dash symbol", "FTH_A_B", "a-b", xtests.MUST_EQUAL),
		xtests.NewCase("dot symbol", "FTH_A__B", "a.b", xtests.MUST_EQUAL),
		xtests.NewCase("dot symbol", "FTH_A___B", "", xtests.MUST_ERROR),
		xtests.NewCase("too many underscroll (env)", "FTH___B", "", xtests.MUST_ERROR),
		xtests.NewCase("too many underscroll (cenv)", "FTC___B", "", xtests.MUST_ERROR),
		xtests.NewCase("too less underscroll (env)", "FTHB", "", xtests.MUST_ERROR),
		xtests.NewCase("too less underscroll (cenv)", "FTCB", "", xtests.MUST_ERROR),
	} {
		assertion.NewName(tc.Name).
			WithActualAndBool(configs.EnvToKey(tc.Actual.(string))).
			WithExpected(tc.Expected).
			Must(tc.Checker...)
	}
}

func TestKeyToEnv(t *testing.T) {
	var assertion = xtests.New(t)
	for _, tc := range []xtests.TestCase{
		xtests.NewCase("1st level key", "test", "FTH_TEST"),
		xtests.NewCase("2nd level key", "test.ntimes", "FTH_TEST__NTIMES"),
		xtests.NewCase("2nd level key with dash", "test.num-ber", "FTH_TEST__NUM_BER"),
		xtests.NewCase("10th level key", "a.b.c.d.e.f.g.h.i.j", "FTH_A__B__C__D__E__F__G__H__I__J"),
		xtests.NewCase("custom key", "_.test", "FTC_TEST"),
		xtests.NewCase("custom key with underscroll", "_.-test", "FTC__TEST"),
	} {
		assertion.NewName(tc.Name).
			WithActual(configs.KeyToEnv(tc.Actual.(string))).
			WithExpected(tc.Expected).
			MustEqual()
	}
}
