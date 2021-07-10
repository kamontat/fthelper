package features

import (
	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/models"
)

func FileSystemOConfig(featureName string, s *fs.ConfigSearcher, def *fs.PathModel) Feature {
	return Raw(featureName, Dependencies{
		KEY_CONFIG: REQUIRE,
	}, func(self *FeatureImpl, input models.Mapper) (interface{}, error) {
		return fs.GetPathModel(GetConfig(input).Data, s, def)
	})
}

func FileSystemConfig(featureName string, s *fs.ConfigSearcher) Feature {
	return FileSystemOConfig(featureName, s, fs.EmptyPathModel)
}
