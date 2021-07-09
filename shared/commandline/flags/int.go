package flags

import (
	"flag"

	"github.com/kamontat/fthelper/shared/maps"
)

type Int struct {
	Name    string
	Default int64
	Usage   string
	Action  func(data int64) maps.Mapper
}

func (f Int) FlagName() string {
	return f.Name
}

func (f Int) Parse(flag *flag.FlagSet) interface{} {
	var result = new(int64)
	flag.Int64Var(result, f.Name, f.Default, f.Usage)
	return result
}

func (f Int) Build(value interface{}) maps.Mapper {
	return f.Action(*value.(*int64))
}
