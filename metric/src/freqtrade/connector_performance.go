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
	var name, expireAt, query = Connector(conn, API_PERF)
	if data, err := conn.Cache(name, expireAt, func() (interface{}, error) {
		var target = make([]*Performance, 0)
		err := conn.GET(name, query, &target)
		return target, err
	}); err == nil && data != nil {
		return data.([]*Performance)
	}

	return EmptyPerformance()
}
