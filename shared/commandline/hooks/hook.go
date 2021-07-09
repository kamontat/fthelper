package hooks

import "github.com/kamontat/fthelper/shared/maps"

// Hook action
type Hook func(config maps.Mapper) error
