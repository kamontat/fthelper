package features

import (
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/models"
)

func Config(config *models.Config) Feature {
	return Raw(KEY_CONFIG, noDeps, withStaticExecutor(config))
}

func AdditionConfig(builder *configs.Builder) Feature {
	return Raw(MergeKeys(KEY_CONFIG, KEY_ADDON), Dependencies{
		KEY_CONFIG:   REQUIRE,
		KEY_STRATEGY: OPTIONAL,
	}, func(self *FeatureImpl, input models.Mapper) (interface{}, error) {
		if input.Has(KEY_STRATEGY) {
			var merger, ok = input.M(KEY_STRATEGY)
			if ok {
				return builder.MergeStrategy(merger).Build(), nil
			}
		}

		return builder.Build(), nil
	})
}
