package runners_test

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/runners"
	"github.com/kamontat/fthelper/shared/xtests"
)

func ToBuffer(summary *runners.Summary) *bytes.Buffer {
	var buffer = new(bytes.Buffer)
	var logger = loggers.Get()

	logger.SetWriter(buffer)
	summary.Log(logger)

	return buffer
}

func TestEmptySummary(t *testing.T) {
	var assertion = xtests.New(t)

	var summary = runners.NewSummary()

	assertion.NewName("get correct name").
		WithExpected(runners.DEFAULT_SUMMARY_NAME).
		WithActual(summary.Name).
		MustEqual()

	var buffer = ToBuffer(summary)
	assertion.NewName("get correct output").
		WithExpected(`
------------------------------------------------------------
ID  Name  Status  Duration
------------------------------------------------------------
Summary: unknown ( 00/00 - NaN%) | 0s

`).
		WithActual(buffer.String()).
		MustEqual()
}

func TestNamedSummary(t *testing.T) {
	var assertion = xtests.New(t)

	var summary = runners.NewNamedSummary("test")

	assertion.NewName("get correct name").
		WithExpected("test").
		WithActual(summary.Name).
		MustEqual()
}

func TestSingleInfoSummary(t *testing.T) {
	var assertion = xtests.New(t)
	var summary = runners.NewSummary().
		Add(runners.NewInformation("info#1").
			SetStatus(runners.INVALID).
			SetDuration(5 * time.Second).SetError(errors.New("test")))

	var buffer = ToBuffer(summary)
	assertion.NewName("get correct output").
		WithExpected(`
------------------------------------------------------------
ID  Name    Status   Duration
1   info#1  invalid  5s
------------------------------------------------------------
Summary: error ( 00/01 - 0.00%) | 5s

`).
		WithActual(buffer.String()).
		MustEqual()
}

func TestMultipleInfoSummary(t *testing.T) {
	var assertion = xtests.New(t)
	var summary = runners.NewSummary().
		Add(runners.NewInformation("info#1").
			SetStatus(runners.SUCCESS).
			SetDuration(5 * time.Second).SetError(errors.New("test"))).
		Add(runners.NewInformation("info#2").
			SetStatus(runners.DISABLED).
			SetDuration(1 * time.Millisecond)).
		Add(runners.NewInformation("info#3").
			SetStatus(runners.SUCCESS).
			SetDuration(791 * time.Microsecond)).
		Add(runners.NewInformation("info#4").
			SetStatus(runners.ERROR).
			SetDuration(23 * time.Millisecond).SetError(errors.New("test")))

	var buffer = ToBuffer(summary)
	assertion.NewName("get correct output").
		WithExpected(`
------------------------------------------------------------
ID  Name    Status    Duration
1   info#1  success   5s
2   info#2  disabled  1ms
3   info#3  success   791µs
4   info#4  error     23ms
------------------------------------------------------------
Summary: partial-success ( 03/04 - 75.00%) | 5.024791s

`).
		WithActual(buffer.String()).
		MustEqual()
}

func TestGroupInfoSummary(t *testing.T) {
	var assertion = xtests.New(t)
	var summary = runners.NewSummary().
		Add(runners.NewInformation("info#01").
			SetStatus(runners.INVALID).
			SetDuration(5*time.Second).SetError(errors.New("test"))).
		Add(runners.NewInformation("info#02").
			SetStatus(runners.DISABLED).
			SetDuration(1*time.Millisecond)).
		AddGroup("test#1",
			runners.NewInformation("info#11").
				SetStatus(runners.DISABLED).
				SetDuration(412*time.Millisecond),
			runners.NewInformation("info#12").
				SetStatus(runners.INVALID).
				SetDuration(13*time.Nanosecond).
				SetError(errors.New("test")),
		).
		AddGroup("test#2",
			runners.NewInformation("info#21").
				SetStatus(runners.SUCCESS).
				SetDuration(213*time.Microsecond),
			runners.NewInformation("info#22").
				SetStatus(runners.SUCCESS).
				SetDuration(999*time.Nanosecond),
		)

	var buffer = ToBuffer(summary)
	assertion.NewName("get correct output").
		WithExpected(`
------------------------------------------------------------
ID      Name     Status    Duration
Group:  default  ------    --------
1       info#01  invalid   5s
2       info#02  disabled  1ms
Group:  test#1   ------    --------
3       info#11  disabled  412ms
4       info#12  invalid   13ns
Group:  test#2   ------    --------
5       info#21  success   213µs
6       info#22  success   999ns
------------------------------------------------------------
Summary: partial-success ( 04/06 - 66.67%) | 5.413214012s

`).
		WithActual(buffer.String()).
		MustEqual()
}
