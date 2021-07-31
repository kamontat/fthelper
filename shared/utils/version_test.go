package utils_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/utils"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestVersion(t *testing.T) {
	var assertion = xtests.New(t)

	var testcases = []xtests.TestCase{
		xtests.NewCase("patch version with prefix", "v0.0.1", float64(1)),
		xtests.NewCase("minor version with prefix", "v0.1.0", float64(100)),
		xtests.NewCase("major version with prefix", "v1.0.0", float64(10000)),
		xtests.NewCase("mixed version with prefix", "v1.2.11", float64(10211)),
		xtests.NewCase("two digit with prefix", "v1.10.10", float64(11010)),
		xtests.NewCase("beta prerelease with prefix", "v1.1.1-beta.1", float64(10101.201)),

		xtests.NewCase("invalid with prefix", "vx.x.x", -1),
		xtests.NewCase("invalid prerelease with prefix", "v1.1.1-test.1", -1),

		xtests.NewCase("patch version", "0.0.1", float64(1)),
		xtests.NewCase("minor version", "0.1.0", float64(100)),
		xtests.NewCase("major version", "1.0.0", float64(10000)),
		xtests.NewCase("mixed version", "1.2.11", float64(10211)),
		xtests.NewCase("two digit", "1.10.10", float64(11010)),
		xtests.NewCase("beta prerelease", "1.1.1-beta.1", float64(10101.201)),

		xtests.NewCase("invalid", "x.x.x", -1),
		xtests.NewCase("invalid prerelease", "1.1.1-test.1", -1),
	}
	for _, testcase := range testcases {
		assertion.NewName(testcase.Name).
			WithActual(utils.VersionNumber(testcase.Actual.(string))).
			WithExpected(testcase.Expected).
			MustEqualFloat()
	}
}
