package runners

import (
	"time"

	"github.com/kamontat/fthelper/shared/errors"
)

type Information interface {
	Name() string
	Status() Status
	Error() *errors.Handler
	Duration() time.Duration

	TotalCount() int
	SuccessCount() int
	FailureCount() int

	String() string
	SString(i time.Time) string
}
