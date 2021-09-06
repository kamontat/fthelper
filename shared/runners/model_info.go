package runners

import (
	"fmt"
	"time"
)

type Status string

const (
	INITIAL  Status = "initial"
	DISABLED Status = "disabled"
	SUCCESS  Status = "success"
	ERROR    Status = "error"
	INVALID  Status = "invalid"
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

func (i *Information) Error() error {
	return i.err
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

// Set first time only, if status was set nothing changes
func (i *Information) SetStatus(status Status) *Information {
	if !i.IsInitial() {
		return i
	}

	i.status = status
	return i
}

func (i *Information) CalDuration(startTime time.Time) *Information {
	i.duration = time.Since(startTime)
	return i
}

func (i *Information) SetDuration(duration time.Duration) *Information {
	i.duration = duration
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
