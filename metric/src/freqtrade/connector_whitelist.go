package freqtrade

type Whitelist struct {
	Length int64
}

func EmptyWhitelist() *Whitelist {
	return &Whitelist{
		Length: 0,
	}
}

func NewWhitelist(conn *Connection) *Whitelist {
	var name, expireAt, query = Connector(conn, API_WHITELIST)
	if data, err := conn.Cache(name, expireAt, func() (interface{}, error) {
		var target = new(Whitelist)
		err := conn.GET(name, query, &target)
		return target, err
	}); err == nil && data != nil {
		return data.(*Whitelist)
	}

	return EmptyWhitelist()
}
