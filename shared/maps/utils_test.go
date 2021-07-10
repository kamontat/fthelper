package maps_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/maps"
)

func TestNormalize(t *testing.T) {
	var i = maps.New().
		Set("$schema", "hello").
		Set("internal.test", true).
		Set("internal.#comment#", "hello").
		Set("b.test", "$schema").
		Set("b.#comment#", maps.New().Set("test", 123)).
		Set("#comment#", maps.New().Set("test", 123))

	var n = maps.Normalize(i.Copy(), []string{"$schema", "#comment#"})

	t.Log(i)
	t.Log(n)
}
