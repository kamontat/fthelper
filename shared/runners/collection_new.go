package runners

func NewCollection(name string, runners ...Runner) *Collection {
	return &Collection{
		Name:    name,
		runners: append([]Runner{}, runners...),
	}
}
