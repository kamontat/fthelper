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

	var queue = datatype.NewLimitQueue(3).
		Enqueue(1).
		Enqueue(2).
		Enqueue(3)

	assertion.NewName("should able to get first queue").
		WithExpected(1).
		WithActual(queue.Get()).
		MustEqual()

	assertion.NewName("should able to remove queue").
		WithExpected(2).
		WithActual(queue.Size()).
		MustEqual()

	var queue1 = datatype.NewLimitQueue(0)

	assertion.NewName("should not able to get first queue").
		WithExpected(nil).
		WithActual(queue1.Get()).
		MustEqual()

	// assertion.NewName("should able to remove queue").
	// 	WithExpected(nil).
	// 	WithActual(queue1.Size()).
	// 	MustEqual()
}
