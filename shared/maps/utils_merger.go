package maps

import (
	"strings"

	"github.com/kamontat/fthelper/shared/utils"
)

// Merge will merge 'a' and 'b'. with optional strategy mapper
// a will be modifiy to be the result
func Merge(a, b map[string]interface{}, strategy Mapper) map[string]interface{} {
	// merge data
	for key, value := range b {
		var replaced = false
		if bData, ok := ToMapper(value); ok {
			if aData, ok := ToMapper(a[key]); ok {
				if exist, ok := strategy.Z(key); ok && exist == MERGER_OVERRIDE {
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
			// delete comment if key prefix and suffix with # { "#foo#": "this is comment" }
			if strings.HasPrefix(key, "#") && strings.HasSuffix(key, "#") {
				delete(a, key)
			} else {
				a[key] = value
			}
		}
	}
	return a
}
