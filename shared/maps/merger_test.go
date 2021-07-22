package maps_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtests"
)

var base = maps.New().
	Set("a", maps.New().
		Set("aa", "aa").
		Set("bb", 15).
		Set("cc", []int{1, 2, 3}),
	).
	Set("b", "b")

var next = maps.New().
	Set("a", maps.New().
		Set("bb", 100).
		Set("cc", []string{"4"}).
		Set("dd", false),
	).
	Set("b", "c").
	Set("c", "d")

func TestMerger(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("normal merge").
		WithExpected(maps.New().
			Set("a", map[string]interface{}{
				"aa": "aa",
				"bb": 100,
				"cc": []interface{}{1, 2, 3, "4"},
				"dd": false,
			}).
			Set("b", "c").
			Set("c", "d")).
		WithActual(maps.Merger(base).Add(next).Merge()).
		MustDeepEqual()

	assertion.NewName("merge with config string").
		WithExpected(maps.New().
			Set("a", maps.New().
				Set("bb", 100).
				Set("cc", []string{"4"}).
				Set("dd", false)).
			Set("b", "c").
			Set("c", "d")).
		WithActual(maps.Merger(base).Add(next).SetConfigValue("a", maps.MERGER_OVERRIDE).Merge()).
		MustDeepEqual()

	assertion.NewName("merge with config").
		WithExpected(maps.New().
			Set("a", map[string]interface{}{
				"aa": "aa",
				"bb": 100,
				"cc": []interface{}{"4"},
				"dd": false,
			}).
			Set("b", "c").
			Set("c", "d")).
		WithActual(maps.Merger(base).Add(next).SetConfig(maps.New().Set("a.cc", maps.MERGER_OVERRIDE)).Merge()).
		MustDeepEqual()
}
