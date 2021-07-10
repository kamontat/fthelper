package xtemplates

import (
	"bytes"
	"text/template"

	"github.com/kamontat/fthelper/shared/utils"
)

func New(name string) *template.Template {
	return template.New(name).
		Funcs(stringFuncs).
		Funcs(jsonFuncs).
		Option("missingkey=error")
}

func File(name string, path string) (*template.Template, error) {
	return New(name).ParseFiles(path)
}

func Buffer(content string, data interface{}, target *bytes.Buffer) error {
	var tpl, err = New(utils.RandString(7)).Parse(content)
	if err != nil {
		return err
	}

	return tpl.Execute(target, data)
}

func Text(content string, data interface{}) (string, error) {
	var target bytes.Buffer
	var err = Buffer(content, data, &target)
	return target.String(), err
}
