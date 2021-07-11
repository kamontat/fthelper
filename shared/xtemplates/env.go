package xtemplates

import (
	"fmt"
	"text/template"

	"github.com/kamontat/fthelper/shared/maps"
)

func byEnv(config maps.Mapper, clusterName string, configPath string) (interface{}, error) {
	var path = fmt.Sprintf("_.%s.%s", clusterName, configPath)
	return config.Get(path)
}

var envFuncs template.FuncMap = map[string]interface{}{
	"byEnv": byEnv,
}
