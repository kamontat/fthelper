package features

type Collection struct {
	features []Feature
}

func (c *Collection) Add(features ...Feature) *Collection {
	c.features = append(c.features, features...)
	return c
}

func (c *Collection) Build(addition ...Feature) []Feature {
	var base []Feature = make([]Feature, 0)
	base = append(base, c.features...)
	base = append(base, addition...)

	return base
}

func NewCollection(features ...Feature) *Collection {
	return &Collection{
		features: features,
	}
}
