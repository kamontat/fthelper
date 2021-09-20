package maps

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/utils"
)

var avoid = []string{"$schema", "#comment#"}

// Merge will merge 'a' and 'b'. with optional strategy mapper
// a will be modifiy to be the result
func Merge(a, b map[string]interface{}, strategy Mapper) map[string]interface{} {
	var log = loggers.Get("map", "merger")

	// merge data
	for key, value := range b {
		var replaced = false
		if bData, ok := ToMapper(value); ok {
			if aData, ok := ToMapper(a[key]); ok {
				if exist, ok := strategy.Z(key); ok && fmt.Sprint(exist) == fmt.Sprint(MERGER_OVERRIDE) {
					log.Debug("found merger type: override (map)")
					a[key] = bData
					replaced = true
				} else {
					a[key] = Merge(aData, bData, strategy.Mi(key))
					replaced = true
				}
			}
		} else if bData, ok := datatype.ToArray(value); ok {
			if aData, ok := datatype.ToArray(a[key]); ok {
				if exist, ok := strategy.Z(key); ok && exist == MERGER_OVERRIDE {
					log.Debug("found merger type: override (array)")
					a[key] = bData
				} else {
					a[key] = append(aData, bData...)
				}

				replaced = true
			}
		}

		if !replaced {
			var oldValue = utils.MaskString(datatype.ForceString(a[key]), utils.MEDIUM)
			var newValue = utils.MaskString(datatype.ForceString(value), utils.MEDIUM)
			// mask map value string
			if oldValue != newValue {
				log.Debug("merging key=%s from %s -> %s",
					key,
					oldValue,
					newValue,
				)
			}

			a[key] = value
		}
	}

	return Normalize(a, avoid)
}
