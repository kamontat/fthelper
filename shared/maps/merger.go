package maps

type MergeConfigType uint8

const (
	MERGER_NORMAL MergeConfigType = iota
	MERGER_OVERRIDE
)

type merger struct {
	Base     Mapper
	addition []Mapper

	Config Mapper
}

func (m *merger) Add(new Mapper) *merger {
	m.addition = append(m.addition, new)
	return m
}

func (m *merger) SetConfig(config Mapper) *merger {
	m.Config = config
	return m
}

func (m *merger) SetConfigValue(key string, value MergeConfigType) *merger {
	m.Config.Set(key, value)
	return m
}

func (m *merger) Merge() Mapper {
	var a = m.Base.Copy()
	for _, data := range m.addition {
		a = Merge(a, data, m.Config)
	}

	return a
}

func Merger(a Mapper) *merger {
	return &merger{
		Base:     a,
		addition: make([]Mapper, 0),
		Config:   New(),
	}
}
