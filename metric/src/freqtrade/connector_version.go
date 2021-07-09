package freqtrade

import "github.com/kamontat/fthelper/shared/maps"

func NewVersion(conn *Connection) string {
	var name, expireAt, query = Connector(conn, API_VERSION)
	if data, err := conn.Cache(name, expireAt, func() (interface{}, error) {
		var target = make(maps.Mapper)
		err := conn.GET(name, query, &target)
		return target.So("version", ""), err
	}); err == nil && data != nil {
		return data.(string)
	}

	return ""
}
