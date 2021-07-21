package maps_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestGetUtilities(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("valid data").
		WithExpected("world").
		WithActualAndError(maps.Get(map[string]interface{}{
			"test": map[string]interface{}{
				"hello": "world",
			},
		}, "test.hello")).
		MustEqual()

	assertion.NewName("valid object data").
		WithExpected(map[string]interface{}{
			"hello": "world",
		}).
		WithActualAndError(maps.Get(map[string]interface{}{
			"test": map[string]interface{}{
				"hello": "world",
			},
		}, "test")).
		MustDeepEqual()

	assertion.NewName("missing data").
		WithExpected(nil).
		WithActualAndError(maps.Get(map[string]interface{}{
			"test": map[string]interface{}{
				"hello": "world",
			},
		}, "test.test")).
		MustEqual()

	assertion.NewName("nil data").
		WithExpected(nil).
		WithActualAndError(maps.Get(map[string]interface{}{
			"test": map[string]interface{}{
				"hello": nil,
			},
		}, "test.hello")).
		MustEqual()
}

func TestGetsUtilities(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("valid data").
		WithExpected(map[string]interface{}{
			"testing": true,
		}).
		WithActualAndError(maps.Gets(map[string]interface{}{
			"test": map[string]interface{}{
				"hello": nil,
				"hello2": map[string]interface{}{
					"testing": true,
				},
			},
		}, "test.hello", "test.hello2")).
		MustDeepEqual()

	assertion.NewName("missing data").
		WithExpected(nil).
		WithActualAndError(maps.Gets(map[string]interface{}{
			"test": map[string]interface{}{
				"hello": nil,
				"hello2": map[string]interface{}{
					"testing": true,
				},
			},
		}, "test.hello", "test.hello3", "test.hello4")).
		MustDeepEqual()
}

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
