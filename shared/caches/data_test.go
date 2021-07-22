package caches_test

import (
	"testing"
	"time"

	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestConstantCacheData(t *testing.T) {
	var assertion = xtests.New(t)
	var normalData = caches.SData("key", "string")

	assertion.NewName("cannot expired").
		WithExpected(true).
		WithActual(normalData.IsExist() && !normalData.IsExpired()).
		MustEqual()

	normalData.Extend()
	assertion.NewName("update updateAt time").
		WithActual(normalData.CreateAt()).
		WithExpected(normalData.UpdateAt()).
		MustNotEqual()

	assertion.NewName("constant data value").
		WithExpected("string").
		WithActual(normalData.Data).
		MustEqual()

	assertion.NewName("not error").
		WithError(normalData.Error).
		MustNotError()
}

func TestDynamicCacheData(t *testing.T) {
	var assertion = xtests.New(t)
	var data = caches.NewData("key", func(o interface{}) (interface{}, error) {
		if o == nil {
			return 0, nil
		}
		return o.(int) + 1, nil
	}, 5*time.Second)

	assertion.NewName("has error by default").
		WithError(data.Error).
		MustError()
	assertion.NewName("data is nil").
		WithActual(data.Data).
		MustBeNil()

	assertion.NewName("after fetch").
		WithDesc("fetch will update if data is missing").
		WithExpected(0).
		WithActualAndError(data.Fetch()).
		Must(xtests.MUST_NOT_ERROR, xtests.MUST_EQUAL)

	assertion.NewName("after 2nd fetch").
		WithDesc("fetch will not update if data is not expired").
		WithExpected(0).
		WithActualAndError(data.Fetch()).
		MustEqual()

	assertion.NewName("after update").
		WithDesc("update should increase value no matter what data is expire or not").
		WithExpected(1).
		WithActualAndError(data.Update()).
		MustEqual()
}
