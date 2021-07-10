package features

import "github.com/kamontat/fthelper/shared/models"

func MergeStrategy(m models.Mapper) Feature {
	return Raw(KEY_STRATEGY, noDeps, withStaticExecutor(m))
}
