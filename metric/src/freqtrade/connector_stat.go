package freqtrade

import "strconv"

type StatObject struct {
	Win  int64 `json:"wins"`
	Loss int64 `json:"losses"`
	Draw int64 `json:"draws"`
}

type DurationObject struct {
	Win  string `json:"wins"`
	Loss string `json:"losses"`
	Draw string `json:"draws"`
}

type Stat struct {
	Reasons  map[string]StatObject `json:"sell_reasons"`
	Duration DurationObject        `json:"durations"`
}

func (s *Stat) toDuration(duration string) float64 {
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

func (s *Stat) WinDuration() float64 {
	return s.toDuration(s.Duration.Win)
}

func (s *Stat) DrawDuration() float64 {
	return s.toDuration(s.Duration.Draw)
}

func (s *Stat) LossDuration() float64 {
	return s.toDuration(s.Duration.Loss)
}

const StatNotANumber = "N/A"

func EmptyStat() *Stat {
	return &Stat{
		Reasons: make(map[string]StatObject),
		Duration: DurationObject{
			Win:  StatNotANumber,
			Loss: StatNotANumber,
			Draw: StatNotANumber,
		},
	}
}

func NewStat(conn *Connection) *Stat {
	var name, expireAt, query = Connector(conn, API_STAT)
	if data, err := conn.Cache(name, expireAt, func() (interface{}, error) {
		var target = new(Stat)
		err := conn.GET(name, query, &target)
		return target, err
	}); err == nil && data != nil {
		return data.(*Stat)
	}

	return EmptyStat()
}
