package features

import (
	"github.com/kamontat/fthelper/shared/models"
	"github.com/kamontat/fthelper/shared/utils"
)

func GetConfig(input models.Mapper) *models.Config {
	var addon = MergeKeys(KEY_CONFIG, KEY_ADDON)

	var raw = input.Ai(KEY_CONFIG).(*models.Config)
	if input.Has(addon) {
		raw = raw.Merge(input.Ai(addon).(*models.Config))
	}
	return raw
}

func MergeKeys(keys ...string) string {
	return utils.JoinString(":", keys...)
}
