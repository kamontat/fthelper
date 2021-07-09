package flags

import (
	"flag"

	"github.com/kamontat/fthelper/shared/maps"
)

type String struct {
	Name    string
	Default string
	Usage   string
	Action  func(data string) maps.Mapper
}

func (f String) FlagName() string {
	return f.Name
}

func (f String) Parse(flag *flag.FlagSet) interface{} {
	var result = new(string)
	flag.StringVar(result, f.Name, f.Default, f.Usage)
	return result
}

func (f String) Build(value interface{}) maps.Mapper {
	return f.Action(*value.(*string))
}
