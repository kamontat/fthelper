package xtemplates_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/xtemplates"
	"github.com/kamontat/fthelper/shared/xtests"
)

func TestJson(t *testing.T) {
	var assertion = xtests.New(t)
	var config = maps.New().
		Set("string", "default").
		Set("int", 12).
		Set("float", 3.99).
		Set("bool", true).
		Set("array", []string{"a", "b", "c"}).
		Set("map", maps.New().
			Set("data", "default").
			Set("float", 9.234))

	assertion.NewName("json").
		WithExpected(`{"array":["a","b","c"],"bool":true,"float":3.99,"int":12,"map":{"data":"default","float":9.234},"string":"default"}`).
		WithActualAndError(xtemplates.Text("{{ json . }}", config)).
		MustEqual()

	assertion.NewName("format json").
		WithExpected(`{
  "array": [
    "a",
    "b",
    "c"
  ],
  "bool": true,
  "float": 3.99,
  "int": 12,
  "map": {
    "data": "default",
    "float": 9.234
  },
  "string": "default"
}`).
		WithActualAndError(xtemplates.Text("{{ indentJson . }}", config)).
		MustEqual()
}
