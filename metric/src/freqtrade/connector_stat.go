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
	if stat, err := FetchStat(conn); err == nil {
		return stat
	}
	return EmptyStat()
}

func FetchStat(conn *Connection) (*Stat, error) {
	var name = API_STAT
	if data, err := conn.Cache(name, conn.ExpireAt(name), func() (interface{}, error) {
		return GetStat(conn)
	}); err == nil {
		return data.(*Stat), nil
	} else {
		return nil, err
	}
}

func GetStat(conn *Connection) (*Stat, error) {
	var target = new(Stat)
	var err = GetConnector(conn, API_STAT, &target)
	return target, err
}
