package features

import "github.com/kamontat/fthelper/shared/models"

func withStaticExecutor(s interface{}) Executor {
	return func(self *FeatureImpl, input models.Mapper) (interface{}, error) {
		// self.Logger.Debug("return static value: %v", s)
		return s, nil
	}
}

func privateWithConfigExecutor(input models.Mapper, configpath string) (interface{}, error) {
	config := GetConfig(input)

	d, err := config.Data.Get(configpath)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func withConfigExecutor(configpath string) Executor {
	return func(self *FeatureImpl, input models.Mapper) (interface{}, error) {
		var data, err = privateWithConfigExecutor(input, configpath)
		if err != nil {
			return nil, err
		} else {
			self.Logger.Debug("receiving '%s' from config path '%s'", data, configpath)
			return data, nil
		}
	}
}

// withOConfigExecutor will get data from config, if any error occurred will return fallback value instead
func withOConfigExecutor(configpath string, def interface{}) Executor {
	return func(self *FeatureImpl, input models.Mapper) (interface{}, error) {
		var data, err = privateWithConfigExecutor(input, configpath)
		if err != nil {
			self.Logger.Debug("error occurred (%s), fallback to %s", err, def)
			return def, nil
		} else {
			self.Logger.Debug("receiving '%s' from config path '%s'", data, configpath)
			return data, nil
		}
	}
}
