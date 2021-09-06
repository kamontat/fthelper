package runners

type Runners struct {
	runners map[string][]*Runner
}

// Add new runner to specify group
func (r *Runners) AddGroup(name string, runners ...*Runner) *Runners {
	if runner, ok := r.runners[name]; ok {
		r.runners[name] = append(runner, runners...)
		return r
	}

	r.runners[name] = runners
	return r
}

// Add new runner to default group
func (r *Runners) Add(runners ...*Runner) *Runners {
	return r.AddGroup(DEFAULT_GROUP_NAME, runners...)
}

func (r *Runners) Run() *Summary {
	var summary = NewSummary()
	for key, runners := range r.runners {
		var informations = make([]*Information, 0)
		for _, runner := range runners {
			informations = append(informations, runner.Run())
		}
		summary.AddGroup(key, informations...)
	}

	return summary
}
