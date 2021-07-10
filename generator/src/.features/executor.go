package features

import "github.com/kamontat/fthelper/shared/models"

type Executor func(self *FeatureImpl, input models.Mapper) (interface{}, error)
