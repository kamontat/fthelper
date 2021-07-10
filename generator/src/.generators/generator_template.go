package generators

import (
	"os"

	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/models"
	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/xtemplates"
)

func Template(name string, fs ...features.Feature) runners.Runner {
	return New(name, features.Dependencies{
		features.KEY_TEMPLATE: features.REQUIRE,
		features.MergeKeys(features.KEY_TEMPLATE, features.KEY_FORMAT): features.OPTIONAL,
		features.KEY_OUTPUT: features.REQUIRE,
		features.MergeKeys(features.KEY_OUTPUT, features.KEY_FORMAT): features.OPTIONAL,
		features.KEY_CONFIG: features.REQUIRE,
		features.MergeKeys(features.KEY_CONFIG, features.KEY_ADDON): features.OPTIONAL,
		features.KEY_KEY:   features.OPTIONAL,
		features.KEY_CHMOD: features.OPTIONAL,
	}, func(param *ExecutorParameter, input models.Mapper) error {
		var config = features.GetConfig(input)
		if input.Has(features.KEY_KEY) {
			var key = input.Ai(features.KEY_KEY).(*configs.KeyModel)
			config = configs.Build(config, key)
		}

		var tplfs = buildFile(features.KEY_TEMPLATE, input)
		param.Logger.Debug("create template: %s", tplfs.Abs())
		var template, err = xtemplates.New(tplfs)
		if err != nil {
			return err
		}

		var outfs = buildFile(features.KEY_OUTPUT, input)
		param.Logger.Debug("create output: %s", outfs.Abs())
		if err := outfs.Build(); err != nil {
			return err
		}
		writer, err := outfs.Writer()
		if err != nil {
			return err
		}

		var addon = fsAddon(config, outfs)
		err = template.Execute(writer, config.MergeMap(addon).Data)
		param.Logger.Info("create output from template at %s", outfs.Abs())
		if a, ok := input.A(features.KEY_CHMOD); err == nil && ok {
			var mode = a.(os.FileMode)
			return writer.Chmod(mode)
		}

		return err
	}, fs...)
}
