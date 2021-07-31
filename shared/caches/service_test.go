package caches_test

import (
	"fmt"
	"testing"

	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestNewCacheService(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("create a cache service").
		WithActual(caches.New()).
		MustNotBeNil()

	assertion.NewName("empty cache size").
		WithExpected(0).
		WithActual(caches.New().Size()).
		MustEqual()

	assertion.NewName("new != global").
		WithExpected(caches.Global).
		WithActual(caches.New()).
		MustNotEqual()

	assertion.NewName("to string").
		WithExpected("{}").
		WithActual(caches.New().String()).
		MustEqual()

	var a = caches.New()
	var err = a.Set("test", true, caches.Persistent)
	assertion.NewName("to string").
		WithExpected(`{"test":{"Data":true}}`).
		WithError(err).
		WithActual(a.String()).
		Must(xtests.MUST_NOT_ERROR, xtests.MUST_EQUAL)
}

func TestCacheService(t *testing.T) {
	var assertion = xtests.New(t)

	var a = caches.New()
	var err = a.Set("hello", "world", "1s")

	assertion.NewName("get size").
		WithExpected(1).
		WithError(err).
		WithActual(a.Size()).
		Must(xtests.MUST_NOT_ERROR, xtests.MUST_EQUAL)

	assertion.NewName("error when empty key data").
		WithExpected("cannot be empty string").
		WithError(a.Set("", true, "1s")).
		MustContainError()

	assertion.NewName("error when existed data").
		WithExpected("use Update() instead").
		WithError(a.SetData(&caches.Data{
			Key:  "hello",
			Data: true,
		})).
		MustContainError()

	assertion.NewName("error when update error").
		WithExpected("cannot do anything").
		WithError(a.SetFn("new", func() (interface{}, error) {
			return nil, fmt.Errorf("cannot do anything")
		}, "1s")).
		MustEqualError()
}
