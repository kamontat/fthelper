package utils_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/utils"
)

func TestString(t *testing.T) {
	t.Run("join partial empty", func(t *testing.T) {
		var expected = "a-b"
		var out = utils.JoinString("-", "a", "", "b")
		if out != expected {
			t.Errorf("%s != %s", out, expected)
		}
	})

	t.Run("join empty", func(t *testing.T) {
		var expected = ""
		var out = utils.JoinString("-", "", "", "")
		if out != expected {
			t.Errorf("%s != %s", out, expected)
		}
	})

	t.Run("join string", func(t *testing.T) {
		var expected = "test.hello.world"
		var out = utils.JoinString(".", "test", "hello", "world")
		if out != expected {
			t.Errorf("%s != %s", out, expected)
		}
	})
}
