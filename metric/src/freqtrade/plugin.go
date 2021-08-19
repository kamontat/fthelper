package freqtrade

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
)

type Plugin interface {
	// Current plugin name, this will be unique
	Name() string

	// Connection is for data from freqtrade itself or from database
	// This function also provide history data from previous build,
	// history data will avaliable only if you add `WithCache()` to the connector
	// otherwise, history data will always be empty.
	Build(connector connection.Connector, connection *connection.Connection, history *datatype.Queue) (interface{}, error)
}
