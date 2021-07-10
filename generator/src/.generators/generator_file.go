package generators

import (
	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/models"
	"github.com/kamontat/fthelper/shared/runners"
)

func CreateFile(name string, fs ...features.Feature) runners.Runner {
	return New(name, features.Dependencies{
		features.KEY_OUTPUT: features.REQUIRE,
		features.MergeKeys(features.KEY_OUTPUT, features.KEY_FORMAT): features.OPTIONAL,
		features.KEY_TEXT: features.OPTIONAL,
	}, func(param *ExecutorParameter, input models.Mapper) error {
		var outfs = buildFile(features.KEY_OUTPUT, input)
		if err := outfs.Build(); err != nil {
			return err
		}

		var text = input.So(features.KEY_TEXT, "")

		param.Logger.Info("create file at %s", outfs.Abs())
		return outfs.Write([]byte(text))
	}, fs...)
}
