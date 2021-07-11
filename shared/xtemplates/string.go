package xtemplates

import (
	"text/template"

	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/utils"
)

func join(input ...interface{}) string {
	var str []string = make([]string, 0)
	for _, i := range input {
		str = append(str, datatype.ForceString(i))
	}

	return utils.JoinString("-", str...)
}

var stringFuncs template.FuncMap = map[string]interface{}{
	"join": join,
}
