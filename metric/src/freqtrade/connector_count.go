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
	if count, err := FetchCount(conn); err == nil {
		return count
	}
	return EmptyCount()
}

func FetchCount(conn *Connection) (*Count, error) {
	var name = API_COUNT
	if data, err := conn.Cache(name, conn.ExpireAt(name), func() (interface{}, error) {
		return GetCount(conn)
	}); err == nil {
		return data.(*Count), nil
	} else {
		return nil, err
	}
}

func GetCount(conn *Connection) (*Count, error) {
	var target = new(Count)
	var err = GetConnector(conn, API_COUNT, &target)
	return target, err
}
