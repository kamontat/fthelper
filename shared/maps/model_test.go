package maps_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/maps"
)

func TestMapper(t *testing.T) {
	t.Run("Create new Mapper", func(t *testing.T) {
		var m = maps.New()
		if m == nil {
			t.Errorf("New() return nil value")
		}
	})
}
