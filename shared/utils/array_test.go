package utils_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/utils"
)

func TestCloneArray(t *testing.T) {
	t.Run("Create immutable array", func(t *testing.T) {
		var arr = []interface{}{"1", "2"}
		var newArr = utils.CloneArray(arr)

		arr = append(arr, "3")
		if len(arr) == len(newArr) {
			t.Errorf("%v should NOT equal to %v", arr, newArr)
		}
	})

	t.Run("Append to new array", func(t *testing.T) {
		var arr = []interface{}{"1", "2"}
		var newArr = utils.CloneArray(arr, "3", "4")

		if len(newArr) != 4 {
			t.Errorf("array should added new element (%v)", newArr)
		}
	})
}

func TestToArray(t *testing.T) {
	t.Run("Create array from interface{}", func(t *testing.T) {
		var inf interface{} = []interface{}{"1", "2"}
		var newArr, ok = utils.ToArray(inf)
		if !ok || len(newArr) != 2 {
			t.Errorf("Interface (%v) should be converted to array", inf)
		}
	})
}
