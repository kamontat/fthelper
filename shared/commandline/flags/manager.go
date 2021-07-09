package flags

import (
	"flag"
	"fmt"

	"github.com/kamontat/fthelper/shared/maps"
)

type Manager struct {
	keys  []string
	flags map[string]Flag
}

func (f *Manager) Add(flag Flag) {
	f.keys = append(f.keys, flag.FlagName())
	f.flags[flag.FlagName()] = flag
}

func (f *Manager) Build(name string, args []string) (map[string]maps.Mapper, []string, error) {
	var option = maps.New()

	// parse option flags
	var flagSet = flag.NewFlagSet(name, flag.ExitOnError)
	for key, builder := range f.flags {
		option.Set(key, builder.Parse(flagSet))
	}
	err := flagSet.Parse(args)
	if err != nil {
		return make(map[string]maps.Mapper), args, err
	}

	// build result from flags
	var arguments = flagSet.Args()
	var result = make(map[string]maps.Mapper)
	for key, value := range option {
		var mapper = f.flags[key].Build(value)
		if !mapper.IsEmpty() {
			result[key] = mapper
		}
	}

	return result, arguments, err
}

func (m *Manager) String() string {
	return fmt.Sprintf("%v", m.keys)
}
