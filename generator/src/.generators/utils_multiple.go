package generators

import (
	"fmt"

	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/models"
	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/utils"
)

// Multiple is deprecated in favor of key-generator
func Multiple(name string, generator func(name string, fs ...features.Feature) runners.Runner, basename string, fs []features.Feature, addons ...features.Feature) *runners.Collection {
	var rs = make([]runners.Runner, 0)
	for n, feature := range fs {
		rs = append(rs, generator(fmt.Sprintf("%s-%d", basename, n), append([]features.Feature{feature}, addons...)...))
	}

	return runners.NewCollection(name, rs...)
}

func MultipleFn(name string, config *models.Config, configpath string, builder func(data interface{}) []runners.Runner) *runners.Collection {
	var rs = make([]runners.Runner, 0)
	var data, err = config.Data.Get(configpath)
	if err == nil {
		var list, ok = utils.ToArray(data)
		if ok {
			for _, d := range list {
				rs = append(rs, builder(d)...)
			}
		}
	}

	return runners.NewCollection(name, rs...)
}
