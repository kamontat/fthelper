package xtemplates_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtemplates"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestXtemplate(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("invalid template").
		WithExpected("function \"invalid\" not defined").
		WithActualAndError(xtemplates.Text("{{ invalid \"function\" }}", maps.New())).
		MustContainError()
}
