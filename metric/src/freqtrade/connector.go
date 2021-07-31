package freqtrade

import (
	"net/url"

	"github.com/kamontat/fthelper/shared/loggers"
)

// Connector is helper to create name, expireAt and url query
func Connector(conn *Connection, name string) (string, string, url.Values) {
	var expireAt = conn.Config.Cache.Get(name)
	var query = conn.Config.Query.Get(name)

	return name, expireAt, query
}

func ConnectorLog(name string) (string, *loggers.Logger) {
	return name, loggers.Get("connector", name)
}

func GetConnector(conn *Connection, name string, target interface{}) error {
	return conn.GET(name, conn.QueryValues(name), target)
}
