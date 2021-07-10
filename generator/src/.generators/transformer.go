package generators

import (
	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/models"
)

var dependenciesKey = "__dependencies__"
var featuresKey = "__features__"

func toMapper(depends features.Dependencies, fs []features.Feature) models.Mapper {
	var mapper = make(models.Mapper)
	for _, f := range fs {
		mapper.Set(f.Name(), f)
	}

	return models.Mapper{
		dependenciesKey: depends,
		featuresKey:     mapper,
	}
}

func toDepends(input interface{}) features.Dependencies {
	var fs, err = input.(models.Mapper).Get(dependenciesKey)
	if err != nil {
		return make(features.Dependencies)
	}

	return fs.(features.Dependencies)
}

func toFeatures(input interface{}) models.Mapper {
	var fs, err = input.(models.Mapper).Get(featuresKey)
	if err != nil {
		return make(models.Mapper)
	}

	return fs.(models.Mapper)
}
