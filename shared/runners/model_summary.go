package runners

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kamontat/fthelper/shared/loggers"
)

const (
	unknownStatus        = "unknown"
	successStatus        = "success"
	errorStatus          = "error"
	partialSuccessStatus = "partial-success"
)

type Summary struct {
	Name         string
	sortKeys     []string
	informations map[string][]*Information
}

func (s *Summary) AddGroup(name string, informations ...*Information) *Summary {
	if info, ok := s.informations[name]; ok {
		s.informations[name] = append(info, informations...)
		return s
	}

	s.sortKeys = append(s.sortKeys, name)
	s.informations[name] = informations
	return s
}

func (s *Summary) Add(informations ...*Information) *Summary {
	return s.AddGroup(DEFAULT_GROUP_NAME, informations...)
}

func (s *Summary) Log(logger *loggers.Logger) {
	var table = logger.Table(4)
	var status = unknownStatus
	var success int = 0
	var total = 0
	var duration time.Duration = 0
	var hasGroup = len(s.sortKeys) > 1

	logger.Newline()
	logger.Line()

	table.Header("ID", "Name", "Status", "Duration")
	for _, key := range s.sortKeys {
		if hasGroup {
			table.Header("Group:", key, "------", "--------")
		}

		for _, information := range s.informations[key] {
			table.Row(
				strconv.Itoa(total+1),
				information.Name(),
				string(information.Status()),
				information.Duration().String(),
			)

			duration += information.Duration()
			total++
			switch information.Status() {
			case SUCCESS, DISABLED:
				success++
				switch status {
				case unknownStatus:
					status = successStatus
				case errorStatus:
					status = partialSuccessStatus
				}
			case ERROR, INVALID, INITIAL:
				switch status {
				case unknownStatus:
					status = errorStatus
				case successStatus:
					status = partialSuccessStatus
				}
			}
		}
	}

	_ = table.End()
	logger.Line()

	logger.Log(fmt.Sprintf(
		"%s: %s ( %02d/%02d - %.2f%%) | %s",
		s.Name,
		status,
		success,
		total,
		(float64(success)*float64(100))/float64(total),
		duration.String(),
	))
	logger.Newline()
}

func NewSummary() *Summary {
	return NewNamedSummary(DEFAULT_SUMMARY_NAME)
}

func NewNamedSummary(name string) *Summary {
	return &Summary{
		Name:         name,
		sortKeys:     make([]string, 0),
		informations: make(map[string][]*Information),
	}
}
