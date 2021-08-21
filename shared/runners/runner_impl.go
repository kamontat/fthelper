package runners

import "time"

// RunnerImpl is implemented of Runner interfaces
type RunnerImpl struct {
	info      *SingleInfo
	validator Validator
	executor  Executor
}

func (r *RunnerImpl) Information() Information {
	return r.info
}

func (r *RunnerImpl) Input(i interface{}) Runner {
	r.info.input = i
	return r
}

func (r *RunnerImpl) updateError(err error) error {
	if r.info.err.And(err).HasError() {
		r.info.status = ERROR
	}

	return err
}

func (r *RunnerImpl) Validate() error {
	if r.info.status == INITIAL {
		var err = r.validator(r.info)
		return r.updateError(err)
	}

	return r.info.err.Error()
}

func (r *RunnerImpl) Run() error {
	if r.info.status != INITIAL {
		return r.info.err.Error()
	}

	startTime := time.Now()
	var err = r.executor(r.info)
	duration := time.Since(startTime)

	r.info.duration = duration

	err = r.updateError(err)
	if err == nil && r.info.status != DISABLED {
		r.info.status = SUCCESS
	}

	return err
}

func NewRunner(name string, vl Validator, ec Executor) Runner {
	return &RunnerImpl{
		validator: vl,
		executor:  ec,
		info:      NewSingleInfo(name),
	}
}
