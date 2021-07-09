package runners

import (
	"fmt"
	"time"

	"github.com/kamontat/fthelper/shared/errors"
)

type SingleInfo struct {
	name     string
	status   Status
	err      *errors.Handler
	duration time.Duration

	input  interface{}
	output interface{}
}

func (i *SingleInfo) Disabled() *SingleInfo {
	i.status = DISABLED
	return i
}

func (i *SingleInfo) In(data interface{}) *SingleInfo {
	i.input = data
	return i
}

func (i *SingleInfo) Input() interface{} {
	return i.input
}

func (i *SingleInfo) Out(data interface{}) *SingleInfo {
	i.output = data
	return i
}

func (i *SingleInfo) Output() interface{} {
	return i.output
}

func (i *SingleInfo) Name() string {
	return i.name
}

func (i *SingleInfo) Status() Status {
	return i.status
}

func (i *SingleInfo) Error() *errors.Handler {
	return i.err
}

func (i *SingleInfo) Duration() time.Duration {
	return i.duration
}

func (i *SingleInfo) TotalCount() int {
	return 1
}

func (i *SingleInfo) SuccessCount() int {
	if i.status == SUCCESS || i.status == DISABLED {
		return 1
	}
	return 0
}

func (i *SingleInfo) FailureCount() int {
	if i.status == INVALID || i.status == ERROR {
		return 1
	}
	return 0
}

func (i *SingleInfo) String() string {
	return fmt.Sprintf("%s: %s | %s", i.Name(), i.Status(), i.Duration().String())
}

func (i *SingleInfo) SString(d time.Time) string {
	return fmt.Sprintf("%s: %s | %s (%s)", i.Name(), i.Status(), i.Duration().String(), time.Since(d).String())
}

func NewSingleInfo(name string) *SingleInfo {
	return &SingleInfo{
		name:     name,
		status:   INITIAL,
		duration: 0,
		err:      errors.New(name),
		input:    nil,
		output:   nil,
	}
}
