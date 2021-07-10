package generators

import (
	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/runners"
)

type Generator func(name string, fs ...features.Feature) runners.Runner
