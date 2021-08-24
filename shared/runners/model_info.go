package runners

import (
	"fmt"
	"time"
)

type Information struct {
	name     string
	status   Status
	err      error
	duration time.Duration
}

func (i *Information) Name() string {
	return i.name
}

func (i *Information) Status() Status {
	return i.status
}

func (i *Information) Run(err error) *Information {
	if !i.IsInitial() {
		return i
	}

	if err != nil {
		i.SetStatus(ERROR)
		i.err = err
	} else {
		i.SetStatus(SUCCESS)
	}
	return i
}

func (i *Information) Duration() time.Duration {
	return i.duration
}

func (i *Information) IsInitial() bool {
	return i.status == INITIAL
}

func (i *Information) SetError(err error) *Information {
	i.err = err
	return i
}

func (i *Information) SetStatus(status Status) *Information {
	if !i.IsInitial() {
		return i
	}

	i.status = status
	return i
}

func (i *Information) SetDuration(startTime time.Time) *Information {
	i.duration = time.Since(startTime)
	return i
}

func (i *Information) String() string {
	return fmt.Sprintf("%s: %s (%v)", i.Name(), i.Status(), i.Duration())
}

func NewInformation(name string) *Information {
	return &Information{
		name:     name,
		status:   INITIAL,
		err:      nil,
		duration: -1,
	}
}
