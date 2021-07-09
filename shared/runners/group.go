package runners

import "time"

// Group is a grouped collection by name
type Group struct {
	Size        int
	naming      []string
	collections []*Collection
}

func (g *Group) New(c *Collection) *Group {
	g.naming = append(g.naming, c.Name)
	g.collections = append(g.collections, c)

	g.Size++
	return g
}

func (g *Group) Run(verbose bool) *Summary {
	startTime := time.Now()
	for _, collection := range g.collections {
		collection.Run()
	}

	info := make([]Information, 0)
	if verbose {
		for _, c := range g.collections {
			info = append(info, c.Informations()...)
		}
		return NewSummary(startTime).As(info...)
	} else {
		for _, c := range g.collections {
			info = append(info, c.Information())
		}
		return NewSummary(startTime).As(info...)
	}
}
