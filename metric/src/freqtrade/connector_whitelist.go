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
	if whitelist, err := FetchWhitelist(conn); err == nil {
		return whitelist
	}
	return EmptyWhitelist()
}

func FetchWhitelist(conn *Connection) (*Whitelist, error) {
	var name = API_WHITELIST
	if data, err := conn.Cache(name, conn.ExpireAt(name), func() (interface{}, error) {
		return GetWhitelist(conn)
	}); err == nil {
		return data.(*Whitelist), nil
	} else {
		return nil, err
	}
}

func GetWhitelist(conn *Connection) (*Whitelist, error) {
	var target = new(Whitelist)
	var err = GetConnector(conn, API_WHITELIST, &target)
	return target, err
}
