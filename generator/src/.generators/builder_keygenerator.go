package generators

import (
	"fmt"

	"github.com/kamontat/fthelper/generators/v4/src/features"
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/runners"
)

type keyGeneratorsBuilder struct {
	name      string
	generator Generator
	keys      []*configs.KeyModel
	features  *features.Collection
}

func (b keyGeneratorsBuilder) Feature(fs ...features.Feature) keyGeneratorsBuilder {
	b.features.Add(fs...)
	return b
}

func (b keyGeneratorsBuilder) FeatureCollection(c *features.Collection) keyGeneratorsBuilder {
	b.features.Add(c.Build()...)
	return b
}

func (b keyGeneratorsBuilder) Build() *runners.Collection {
	var rs = make([]runners.Runner, 0)
	for _, key := range b.keys {
		var gen = b.generator(
			fmt.Sprintf("%s (%s)", b.name, key.String()),
			b.features.Build(features.KeyModel(key))...,
		)

		rs = append(rs, gen)
	}

	return runners.NewCollection(
		b.name,
		rs...,
	)
}

func NewKeyGeneratorBuilder(name string, generator Generator, keys ...*configs.KeyModel) keyGeneratorsBuilder {
	return keyGeneratorsBuilder{
		name:      name,
		generator: generator,
		keys:      keys,
		features:  features.NewCollection(),
	}
}
