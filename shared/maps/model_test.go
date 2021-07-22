package maps_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestNewMapper(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("new mapper").
		WithActual(maps.New()).
		MustNotBeNil()
	assertion.NewName("new mapper length").
		WithExpected(0).
		WithActual(len(maps.New())).
		MustEqual()
	assertion.NewName("new mapper is empty").
		WithExpected(true).
		WithActual(maps.New().IsEmpty()).
		MustEqual()
}

func TestSetMapper(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("set mapper").
		WithExpected(maps.Mapper(map[string]interface{}{
			"a": true,
		})).
		WithActual(maps.New().Set("a", true)).
		MustDeepEqual()

	assertion.NewName("set recursive mapper").
		WithExpected(maps.Mapper(map[string]interface{}{
			"a": map[string]interface{}{
				"b": true,
			},
		})).
		WithActual(maps.New().Set("a.b", true)).
		MustDeepEqual()
}

func TestGetMapper(t *testing.T) {
	var assertion = xtests.New(t)

	var mapper = maps.New().
		Set("a", "ant").
		Set("b", 199).
		Set("c", true).
		Set("d", 0.002).
		Set("e", maps.New().
			Set("ea", "eat").
			Set("eb", 299).
			Set("ec", false).
			Set("ed", []string{"a", "b", "c"})).
		Set("f", []int{9, 8, 7})

	assertion.NewName("get root value").
		WithExpected(int64(199)).
		WithActualAndError(mapper.Ie("b")).
		MustEqual()

	assertion.NewName("get array value").
		WithExpected([]interface{}{9, 8, 7}).
		WithActualAndError(mapper.Ae("f")).
		MustDeepEqual()

	assertion.NewName("error on not exist value").
		WithActualAndError(mapper.Ze("not-exist-value")).
		MustError()
}
