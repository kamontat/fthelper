package generators

import (
	"fmt"

	"github.com/kamontat/fthelper/generator/v4/src/generators/plugins"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func GetRunner(data maps.Mapper, config maps.Mapper) (runners.Runner, error) {
	var name = data.Si("type")
	switch name {
	case "json":
		return plugins.Json(data, config), nil
	case "create":
		return plugins.Create(data, config), nil
	case "copy":
		return plugins.Copy(data, config), nil
	case "template":
		return plugins.Template(data, config), nil
	case "strategy":
		return plugins.CStrategy(data, config), nil
	case "config":
		return plugins.CConfig(data, config), nil
	}

	return nil, fmt.Errorf("cannot found generator for type '%s'", name)
}
