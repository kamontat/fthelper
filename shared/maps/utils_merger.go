package maps

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/utils"
)

var avoid = []string{"$schema", "#comment#"}

// Merge will merge 'a' and 'b'. with optional strategy mapper
// a will be modifiy to be the result
func Merge(a, b map[string]interface{}, strategy Mapper) map[string]interface{} {
	// merge data
	for key, value := range b {
		var replaced = false
		if bData, ok := ToMapper(value); ok {
			if aData, ok := ToMapper(a[key]); ok {
				if exist, ok := strategy.Z(key); ok && fmt.Sprint(exist) == fmt.Sprint(MERGER_OVERRIDE) {
					a[key] = bData
					replaced = true
				} else {
					a[key] = Merge(aData, bData, strategy.Mi(key))
					replaced = true
				}
			}
		} else if bData, ok := utils.ToArray(value); ok {
			if aData, ok := utils.ToArray(a[key]); ok {
				if exist, ok := strategy.Z(key); ok && exist == MERGER_OVERRIDE {
					a[key] = bData
				} else {
					a[key] = append(aData, bData...)
				}

				replaced = true
			}
		}

		if !replaced {
			a[key] = value
		}
	}

	return Normalize(a, avoid)
}
