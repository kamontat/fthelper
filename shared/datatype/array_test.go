package datatype_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/datatype"
)

func TestToArray(t *testing.T) {
	t.Run("Create array from interface{}", func(t *testing.T) {
		var inf interface{} = []interface{}{"1", "2"}
		var newArr, ok = datatype.ToArray(inf)
		if !ok || len(newArr) != 2 {
			t.Errorf("Interface (%v) should be converted to array", inf)
		}
	})
}
