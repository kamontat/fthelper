package freqtrade

type Count struct {
	Current int64
	Max     int64

	TotalStake float64 `json:"total_stake"`
}

func EmptyCount() *Count {
	return &Count{
		Current:    -1,
		Max:        -1,
		TotalStake: 0,
	}
}

func NewCount(conn *Connection) *Count {
	var name, expireAt, query = Connector(conn, API_COUNT)
	if data, err := conn.Cache(name, expireAt, func() (interface{}, error) {
		var target = new(Count)
		err := conn.GET(name, query, &target)
		return target, err
	}); err == nil && data != nil {
		return data.(*Count)
	}

	return EmptyCount()
}
