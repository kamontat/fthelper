package runners

import "time"

func NewSummary(startTime time.Time) *Summary {
	return &Summary{
		sorting:     make([]string, 0),
		information: make(map[string]Information),
		startTime:   startTime,
	}
}

func RunSummary(r Runner) *Summary {
	startTime := time.Now()

        _ = r.Validate()
	_ = r.Run()

	return NewSummary(startTime).A(r.Information())
}

func ColSummary(c *Collection) *Summary {
	startTime := time.Now()

	c.Run()

	return NewSummary(startTime).As(c.Informations()...)
}
