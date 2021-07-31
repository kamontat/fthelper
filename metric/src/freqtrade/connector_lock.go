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
	if locks, err := FetchLocks(conn); err == nil {
		return locks
	}
	return EmptyLocks()
}

func FetchLocks(conn *Connection) (*Locks, error) {
	var name = API_LOCK
	if data, err := conn.Cache(name, conn.ExpireAt(name), func() (interface{}, error) {
		return GetLocks(conn)
	}); err == nil {
		return data.(*Locks), nil
	} else {
		return nil, err
	}
}

func GetLocks(conn *Connection) (*Locks, error) {
	var target = new(Locks)
	var err = GetConnector(conn, API_LOCK, &target)
	return target, err
}
