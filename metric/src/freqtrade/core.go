package freqtrade

import "github.com/kamontat/fthelper/metric/v4/src/connection"

func Build(connections []*connection.Connection) (result []connection.Connector) {
	result = make([]connection.Connector, 0)
	for _, connection := range connections {
		result = append(result, New(connection))
	}
	return
}
