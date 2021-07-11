package xtemplates

import (
	"strconv"
	"strings"
	"text/template"
)

func ratio(percentString string) (float64, error) {
	percent, err := strconv.ParseFloat(strings.TrimSuffix(percentString, "%"), 64)
	if err != nil {
		return -1, err
	}

	return percent / 100, nil
}

var numberFuncs template.FuncMap = map[string]interface{}{
	"ratio": ratio,
}
