package freqtrade

func NewPing(conn *Connection) bool {
	var name, expireAt, query = Connector(conn, API_PING)
	if data, err := conn.Cache(name, expireAt, func() (interface{}, error) {
		var target = make(map[string]interface{})
		err := conn.GET(name, query, &target)
		return target["status"] == "pong", err
	}); err == nil && data != nil {
		return data.(bool)
	}

	return false
}

func NewPingI(conn *Connection) int {
	var b = NewPing(conn)
	if b {
		return 1
	} else {
		return 0
	}
}
