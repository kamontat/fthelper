package generators

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func GetRunner(data maps.Mapper, fsConfig maps.Mapper) (runners.Runner, error) {
	var name = data.Si("type")
	switch name {
	case "json":
		// add correctly generator
		return Json(data, fsConfig), nil
	}

	return nil, fmt.Errorf("cannot found generator for type %s", name)
}
