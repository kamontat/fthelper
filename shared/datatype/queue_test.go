package datatype_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestNormalQueue(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("create normal queue").
		WithExpected(0).
		WithActual(datatype.NewQueue().Size()).
		MustEqual()

	assertion.NewName("create limit queue").
		WithExpected(2).
		WithActual(datatype.
			NewLimitQueue(2).
			Enqueue(1).
			Enqueue(2).
			Enqueue(3).
			Enqueue(4).
			Enqueue(5).
			Size()).
		MustEqual()

	assertion.NewName("correct head on limit queue").
		WithExpected(3).
		WithActual(datatype.
			NewLimitQueue(3).
			Enqueue(1).
			Enqueue(2).
			Enqueue(3).
			Enqueue(4).
			Enqueue(5).
			Head()).
		MustEqual()

	assertion.NewName("correct tail on limit queue").
		WithExpected(5).
		WithActual(datatype.
			NewLimitQueue(3).
			Enqueue(1).
			Enqueue(2).
			Enqueue(3).
			Enqueue(4).
			Enqueue(5).
			Tail()).
		MustEqual()
}
