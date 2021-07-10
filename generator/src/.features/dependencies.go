package features

type DependType string

func (t DependType) IsRequire() bool {
	return t == REQUIRE
}

func (t DependType) IsOptional() bool {
	return t == OPTIONAL
}

const (
	REQUIRE  DependType = "require"
	OPTIONAL DependType = "optional"
)

type Dependencies map[string]DependType

func (d Dependencies) Keys() []string {
	var arr = make([]string, 0)
	for key := range d {
		arr = append(arr, key)
	}

	return arr
}

func (d Dependencies) Has(key string) bool {
	_, ok := d[key]
	return ok
}
