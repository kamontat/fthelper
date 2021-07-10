package generators

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/models"
	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/utils"
	"github.com/kamontat/fthelper/shared/xtemplates"
)

func Json(name string, fs ...features.Feature) runners.Runner {
	return New(name, features.Dependencies{
		features.KEY_TEMPLATE: features.REQUIRE,
		features.MergeKeys(features.KEY_TEMPLATE, features.KEY_FORMAT): features.OPTIONAL,
		features.KEY_OUTPUT: features.REQUIRE,
		features.MergeKeys(features.KEY_OUTPUT, features.KEY_FORMAT): features.OPTIONAL,
		features.KEY_CONFIG: features.REQUIRE,
		features.MergeKeys(features.KEY_CONFIG, features.KEY_ADDON): features.OPTIONAL,
		features.KEY_KEY:      features.OPTIONAL,
		features.KEY_STRATEGY: features.OPTIONAL,
	}, func(param *ExecutorParameter, input models.Mapper) error {
		var key = input.Ai(features.KEY_KEY).(*configs.KeyModel)
		var config = configs.Build(features.GetConfig(input), key)

		// build output file
		var outfs = buildFile(features.KEY_OUTPUT, input)
		if err := outfs.Build(); err != nil {
			return err
		}
		// add addon data to config
		var addon = fsAddon(config, outfs)
		config = config.MergeMap(addon)

		// get json merge strategy
		var merger = input.Mi(features.KEY_STRATEGY)

		var base models.Mapper = make(map[string]interface{})
		var files = buildFiles(features.KEY_TEMPLATE, input)
		for _, file := range files {
			var template, err = xtemplates.New(file)
			if err != nil {
				return err
			}

			var content bytes.Buffer
			err = template.Execute(&content, config.Data)
			if err != nil {
				return err
			}

			var current map[string]interface{}
			err = json.Unmarshal(content.Bytes(), &current)
			if err != nil {
				return fmt.Errorf("invalid json string: %s (%s)", content.String(), err.Error())
			}

			base = utils.MergeJson(base, current, merger)
		}

		// build json from interface{}
		dataBytes, err := json.MarshalIndent(base, "", "  ") // add override data if exist
		if err != nil {
			return err
		}

		param.Logger.Info("merge json template to %s", outfs.Abs())
		return outfs.Write(dataBytes)
	}, fs...)
}
