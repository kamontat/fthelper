package freqtrade

type Performance struct {
	Pair      string  `json:"pair"`
	Profit    float64 `json:"profit"`
	ProfitAbs float64 `json:"profit_abs"`
	Count     int64   `json:"count"`
}

func EmptyPerformance() []*Performance {
	return make([]*Performance, 0)
}

func NewPerformance(conn *Connection) []*Performance {
	if perf, err := FetchPerformance(conn); err == nil {
		return perf
	}
	return EmptyPerformance()
}

func FetchPerformance(conn *Connection) ([]*Performance, error) {
	var name = API_PERF
	if data, err := conn.Cache(name, conn.ExpireAt(name), func() (interface{}, error) {
		return GetPerformance(conn)
	}); err == nil {
		return data.([]*Performance), nil
	} else {
		return nil, err
	}
}

func GetPerformance(conn *Connection) ([]*Performance, error) {
	var target = make([]*Performance, 0)
	var err = GetConnector(conn, API_PERF, &target)
	return target, err
}
