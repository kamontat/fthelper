package flags

import (
	"flag"

	"github.com/kamontat/fthelper/shared/maps"
)

type Flag interface {
	FlagName() string
	Parse(flag *flag.FlagSet) interface{}
	Build(value interface{}) maps.Mapper
}
