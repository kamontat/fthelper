package runners

import (
	"github.com/kamontat/fthelper/shared/errors"
)

// Collection is multiple runners
type Collection struct {
	Name    string
	runners []Runner
}

func (c *Collection) Runners() []Runner {
	return c.runners
}

func (c *Collection) Add(r Runner) *Collection {
	c.runners = append(c.runners, r)
	return c
}

func (c *Collection) Merge(cc *Collection) *Collection {
	c.runners = append(c.runners, cc.runners...)
	return c
}

func (c *Collection) Run() *errors.Handler {
	var errs = errors.New()
	for _, runner := range c.runners {
		if err := runner.Validate(); err != nil {
			errs.And(err)
			continue
		}

		if err := runner.Run(); err != nil {
			errs.And(err)
		}
	}

	return errs
}

// Informations return information of runners
func (c *Collection) Informations() []Information {
	var info = make([]Information, len(c.runners))
	for i, r := range c.runners {
		info[i] = r.Information()
	}

	return info
}

func (c *Collection) Information() Information {
	var infos = c.Informations()
	return NewMultipleNamedInfo(c.Name, infos...)
}
