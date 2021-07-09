package freqtrade

type Lock struct {
	Id             int
	Active         bool
	StartTimestamp int64 `json:"lock_end_timestamp"`
	EndTimestamp   int64 `json:"lock_timestamp"`
	Pair           string
	Reason         string
}

type Locks struct {
	Count int    `json:"lock_count"`
	List  []Lock `json:"locks"`
}

func EmptyLocks() *Locks {
	return &Locks{
		Count: 0,
		List:  make([]Lock, 0),
	}
}

func NewLocks(conn *Connection) *Locks {
	var name, expireAt, query = Connector(conn, API_LOCK)
	if data, err := conn.Cache(name, expireAt, func() (interface{}, error) {
		var target = new(Locks)
		err := conn.GET(name, query, &target)
		return target, err
	}); err == nil && data != nil {
		return data.(*Locks)
	}

	return EmptyLocks()
}
