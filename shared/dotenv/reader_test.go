package dotenv_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/dotenv"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestReader(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("read normal environment format").
		WithActualAndError(dotenv.Unmarshal("test=true")).
		WithExpected(map[string]string{
			"test": "true",
		}).
		MustDeepEqual()

	assertion.NewName("read normal environment format").
		WithActualAndError(dotenv.Unmarshal("# test=true")).
		WithExpected(map[string]string{}).
		Must(xtests.MUST_NOT_ERROR, xtests.MUST_DEEP_EQUAL)
}
