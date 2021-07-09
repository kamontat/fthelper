package runners

import (
	"fmt"
	"strings"
	"time"

	"github.com/kamontat/fthelper/shared/errors"
)

type MultipleInfo struct {
	name     string
	errs     *errors.Handler
	duration time.Duration

	total   int
	success int
	failure int

	info []Information
}

func (i *MultipleInfo) Name() string {
	if i.TotalCount() == 1 {
		return i.info[0].Name()
	} else {
		var prefix = ""
		if i.name != "" {
			prefix = i.name + " "
		}

		var name = fmt.Sprintf("%d tasks", i.TotalCount())
		return fmt.Sprintf("%s%s", prefix, name)
	}

}

func (i *MultipleInfo) Status() Status {
	if i.SuccessCount() == 0 {
		return ERROR
	} else if i.SuccessCount() < i.TotalCount() {
		return PARTIAL_SUCCESS
	} else {
		return SUCCESS
	}
}

func (i *MultipleInfo) Error() *errors.Handler {
	return i.errs
}

func (i *MultipleInfo) Duration() time.Duration {
	return i.duration
}

func (i *MultipleInfo) TotalCount() int {
	return i.total
}

func (i *MultipleInfo) SuccessCount() int {
	return i.success
}

func (i *MultipleInfo) FailureCount() int {
	return i.failure
}

func (i *MultipleInfo) String() string {
	var total = i.TotalCount()
	var success = i.SuccessCount()
	return fmt.Sprintf("%s: %s ( %0d/%02d - %.2f%%) | %s",
		strings.ToTitle(i.Name()),
		i.Status(),
		success,
		total,
		(float64(success)*float64(100))/float64(total),
		i.duration.String(),
	)
}

func (i *MultipleInfo) SString(d time.Time) string {
	var total = i.TotalCount()
	var success = i.SuccessCount()
	return fmt.Sprintf("%s: %s ( %0d/%02d - %.2f%%) | %s (%s)",
		strings.ToTitle(i.Name()),
		i.Status(),
		success,
		total,
		(float64(success)*float64(100))/float64(total),
		i.duration.String(),
		time.Since(d).String(),
	)
}

func NewMultipleNamedInfo(name string, info ...Information) *MultipleInfo {
	var errBase = errors.New("informations", name)
	var baseDuration time.Duration = 0

	var total = 0
	var success = 0
	var failure = 0

	for _, i := range info {
		errBase.Merge(i.Error())
		baseDuration += i.Duration()

		total += i.TotalCount()
		success += i.SuccessCount()
		failure += i.FailureCount()
	}

	return &MultipleInfo{
		name:     name,
		info:     info,
		errs:     errBase,
		duration: baseDuration,

		total:   total,
		success: success,
		failure: failure,
	}
}

func NewMultipleInfo(info ...Information) *MultipleInfo {
	return NewMultipleNamedInfo("", info...)
}
