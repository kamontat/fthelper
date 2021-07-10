package xtemplates

import (
	"text/template"

	"github.com/kamontat/fthelper/shared/maps"
)

func json(m maps.Mapper) (string, error) {
	var a, e = maps.ToJson(m)
	return string(a), e
}

var jsonFuncs template.FuncMap = map[string]interface{}{
	"json": json,
}
