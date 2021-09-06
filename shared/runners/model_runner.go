package runners

import "time"

type Runner struct {
	Name      string
	Context   *Context
	validator Validator
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

	// Check `Disabled` before user .Run() runner
	if r.Context.Disabled {
		return information.
			CalDuration(startTime).
			SetStatus(DISABLED)
	}

	var err = r.validator(r.Context)
	// Check `Disabled` status again after run validator
	if r.Context.Disabled {
		return information.
			CalDuration(startTime).
			SetStatus(DISABLED)
	}

	if err != nil {
		return information.
			SetError(err).
			CalDuration(startTime).
			SetStatus(INVALID)
	}

	err = r.executor(r.Context)
	if err != nil {
		return information.
			SetError(err).
			CalDuration(startTime).
			SetStatus(ERROR)
	}

	return information.
		CalDuration(startTime).
		SetStatus(SUCCESS)
}

func NewRunner(name string, validator Validator, executor Executor) *Runner {
	return &Runner{
		Name:      name,
		Context:   NewContext(name),
		validator: validator,
		executor:  executor,
	}
}
