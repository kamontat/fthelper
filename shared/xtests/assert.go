package xtests

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/kamontat/fthelper/shared/datatype"
)

type Assertion struct {
	name     string
	actual   interface{}
	expected interface{}
	err      error

	T *testing.T
}

func (a *Assertion) WithExpected(e interface{}) *Assertion {
	a.expected = e
	return a
}

func (a *Assertion) WithActual(v interface{}) *Assertion {
	a.actual = v
	return a
}

func (a *Assertion) WithActualAndError(v interface{}, err error) *Assertion {
	a.actual = v
	a.err = err
	return a
}

func (a *Assertion) WithActualAndBool(v interface{}, ok bool) *Assertion {
	if ok {
		a.actual = v
	} else {
		a.err = fmt.Errorf("result (%v) is not ok", v)
	}
	return a
}

func (a *Assertion) WithError(err error) *Assertion {
	a.err = err
	return a
}

func (a *Assertion) must(fn func(t *testing.T) bool) bool {
	if a.name != "" {
		return a.T.Run(a.name, func(t *testing.T) {
			fn(t)
		})
	}
	return fn(a.T)
}
func (a *Assertion) mustBoolean(valid bool, msg string, param ...interface{}) bool {
	return a.must(func(t *testing.T) bool {
		if !valid {
			t.Errorf(msg, param...)
		}
		return valid
	})
}

func (a *Assertion) MustNotError() bool {
	return a.mustBoolean(a.err == nil, "we found error: %v", a.err)
}

func (a *Assertion) MustError() bool {
	return a.mustBoolean(a.err != nil, "we expected some error, but nil is return")
}

func (a *Assertion) MustEqual() bool {
	return a.mustBoolean(a.actual == a.expected, "we expected '%v', but got '%v' instead", a.expected, a.actual)
}

func (a *Assertion) MustDeepEqual() bool {
	return a.mustBoolean(reflect.DeepEqual(a.actual, a.expected), "we expected '%v', but got '%v' instead (as deep-equal)", a.expected, a.actual)
}

func (a *Assertion) MustEqualString() bool {
	var actual = datatype.ForceString(a.actual)
	var expected = datatype.ForceString(a.expected)
	return a.mustBoolean(actual == expected, "we expected '%v', but got '%v' instead (as string)", a.expected, a.actual)
}

func (a *Assertion) MustNotEqual() bool {
	return a.mustBoolean(a.actual != a.expected, "we expected '%v' and '%v' to be different", a.expected, a.actual)
}

func (a *Assertion) MustEqualError() bool {
	return a.mustBoolean(a.err.Error() == a.expected, "we expected '%v' error, but got '%v' error instead", a.expected, a.err)
}

func (a *Assertion) MustContainError() bool {
	var expected = datatype.ForceString(a.expected)
	return a.mustBoolean(strings.Contains(a.err.Error(), expected), "we expected error should contains '%v', but got '%v' instead", expected, a.err)
}

func (a *Assertion) Must(checker ...MustChecker) bool {
	var mapper map[MustChecker]func() bool = make(map[MustChecker]func() bool)

	mapper[MUST_ERROR] = a.MustError
	mapper[MUST_NOT_ERROR] = a.MustNotError
	mapper[MUST_EQUAL] = a.MustEqual
	mapper[MUST_DEEP_EQUAL] = a.MustDeepEqual
	mapper[MUST_EQUAL_STRING] = a.MustEqualString
	mapper[MUST_NOT_EQUAL] = a.MustNotEqual
	mapper[MUST_EQUAL_ERROR] = a.MustEqualError
	mapper[MUST_CONTAINS_ERROR] = a.MustContainError

	for _, c := range checker {
		if fn, ok := mapper[c]; ok {
			if !fn() {
				return false
			}
		} else {
			return a.must(func(t *testing.T) bool {
				t.Errorf("checker %s type is not valid type", c)
				return false
			})
		}
	}

	return true
}
