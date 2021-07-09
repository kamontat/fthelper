package runners

import (
	"strconv"
	"time"

	"github.com/kamontat/fthelper/shared/loggers"
)

type Summary struct {
	sorting     []string // sorting is a array of information name
	information map[string]Information
	startTime   time.Time
}

func (s *Summary) Add(name string, i Information) *Summary {
	s.sorting = append(s.sorting, name)
	s.information[name] = i
	return s
}

func (s *Summary) A(i Information) *Summary {
	return s.Add(i.Name(), i)
}

func (s *Summary) As(is ...Information) *Summary {
	for _, i := range is {
		s.A(i)
	}

	return s
}

func (s *Summary) Log(logger *loggers.Logger) {
	var informations []Information = make([]Information, 0)

	logger.Newline()
	logger.Line()
	table := logger.Table(4)
	table.Header("ID", "Name", "Status", "Duration")
	for i, name := range s.sorting {
		var info = s.information[name]

		informations = append(informations, info)
		table.Row(
			strconv.Itoa(i),
			info.Name(),
			string(info.Status()),
			info.Duration().String(),
		)
	}
	table.End()
	logger.Line()

	var aggregator = NewMultipleNamedInfo("summary", informations...)
	logger.Log(aggregator.SString(s.startTime))
	logger.Line()
	logger.Newline()

	logger.ErrorExit(aggregator.Error())
}
