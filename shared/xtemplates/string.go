package xtemplates

import (
	"text/template"

	"github.com/kamontat/fthelper/shared/utils"
)

func join(str ...string) string {
	return utils.JoinString("-", str...)
}

var stringFuncs template.FuncMap = map[string]interface{}{
	"join": join,
}
