package freqtrade

import "github.com/kamontat/fthelper/shared/maps"

func NewVersion(conn *Connection) string {
	var version, _ = FetchVersion(conn)
	return version
}

func FetchVersion(conn *Connection) (string, error) {
	var name = API_VERSION
	if data, err := conn.Cache(name, conn.ExpireAt(name), func() (interface{}, error) {
		return GetVersion(conn)
	}); err == nil {
		return data.(string), nil
	} else {
		return "", err
	}
}

func GetVersion(conn *Connection) (string, error) {
	var target = make(maps.Mapper)
	var err = GetConnector(conn, API_VERSION, &target)
	if err != nil {
		return "", err
	}
	return target.Se("version")
}
