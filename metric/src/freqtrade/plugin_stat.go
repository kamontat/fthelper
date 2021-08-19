package freqtrade

import (
	"strconv"

	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
)

const (
	STAT_CONST     = "stats"
	StatNotANumber = "N/A"
)

func NewStat() *stat {
	return &stat{
		Reasons: make(map[string]statObject),
		Duration: durationObject{
			Win:  StatNotANumber,
			Loss: StatNotANumber,
			Draw: StatNotANumber,
		},
	}
}

type statObject struct {
	Win  int64 `json:"wins"`
	Loss int64 `json:"losses"`
	Draw int64 `json:"draws"`
}

type durationObject struct {
	Win  string `json:"wins"`
	Loss string `json:"losses"`
	Draw string `json:"draws"`
}

type stat struct {
	Reasons  map[string]statObject `json:"sell_reasons"`
	Duration durationObject        `json:"durations"`
}

func (s *stat) toDuration(duration string) float64 {
	const defaultValue = -1
	if duration == StatNotANumber {
		return defaultValue
	}

	f, e := strconv.ParseFloat(duration, 64)
	if e != nil {
		return defaultValue
	}
	return f
}

func (s *stat) WinDuration() float64 {
	return s.toDuration(s.Duration.Win)
}

func (s *stat) DrawDuration() float64 {
	return s.toDuration(s.Duration.Draw)
}

func (s *stat) LossDuration() float64 {
	return s.toDuration(s.Duration.Loss)
}

func (s *stat) Name() string {
	return STAT_CONST
}

func (s *stat) Build(connector connection.Connector, connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	err := connection.Http.GET(s.Name(), s)
	return s, err
}

func ToStat(connector connection.Connector) (*stat, error) {
	raw, err := connector.Connect(STAT_CONST)
	if err != nil {
		return nil, err
	}
	return raw.(*stat), nil
}
