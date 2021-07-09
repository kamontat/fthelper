package runners

func NewGroup() *Group {
	return &Group{
		Size:        0,
		naming:      make([]string, 0),
		collections: make([]*Collection, 0),
	}
}
