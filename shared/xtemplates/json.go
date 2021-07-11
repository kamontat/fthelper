package xtemplates

import (
	"text/template"

	"github.com/kamontat/fthelper/shared/maps"
)

func json(m interface{}) (string, error) {
	var a, e = maps.ToJson(m)
	return string(a), e
}

func indentJson(m interface{}) (string, error) {
	var a, e = maps.ToFormatJson(m)
	return string(a), e
}

var jsonFuncs template.FuncMap = map[string]interface{}{
	"json":       json,
	"indentJson": indentJson,
}
