package runners

import "time"

type Runner struct {
	Name      string
	Context   *Context
	validator Executor
	executor  Executor
}

func (r *Runner) Input(input interface{}) *Runner {
	r.Context.In(input)
	return r
}

func (r *Runner) Disable(disable bool) *Runner {
	r.Context.Disabled = disable
	return r
}

func (r *Runner) Run() *Information {
	var information = NewInformation(r.Name)
	var startTime = time.Now()

	if r.Context.Disabled {
		return information.
			SetDuration(startTime).
			SetStatus(DISABLED)
	}

	var err = r.validator(r.Context)
	if err != nil {
		return information.
			SetError(err).
			SetDuration(startTime).
			SetStatus(INVALID)
	}

	err = r.executor(r.Context)
	if err != nil {
		return information.
			SetError(err).
			SetDuration(startTime).
			SetStatus(ERROR)
	}

	return information.
		SetDuration(startTime).
		SetStatus(SUCCESS)
}

func NewRunner(name string, validator, executor Executor) *Runner {
	return &Runner{
		Name:      name,
		Context:   NewContext(name),
		validator: validator,
		executor:  executor,
	}
}
