package connection

import (
	"fmt"

	"github.com/kamontat/fthelper/metric/v4/src/clients"
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/maps"
)

type Connection struct {
	Cluster string
	Http    *clients.Http
	Db      *clients.Database
}

func (c *Connection) String() string {
	return fmt.Sprintf(`Connection cluster '%s'
  http     - %v
  database - %v`, c.Cluster, c.Http, c.Db)
}

func NewConnections(config maps.Mapper) ([]*Connection, error) {
	var clusters = config.Ao("clusters", []interface{}{""})

	var connections = make([]*Connection, 0)
	for _, raw := range clusters {
		var cluster = raw.(string)

		var finalConfig = configs.BuildClusterConfig(cluster, config)
		var freqtradeConfig = finalConfig.Mi("freqtrade")

		http, err := clients.NewHttp(cluster, freqtradeConfig.Mi("http"))
		if err != nil {
			return connections, err
		}

		db, err := clients.NewDatabase(cluster, freqtradeConfig.Mi("db"))
		if err != nil {
			return connections, err
		}

		connections = append(connections, &Connection{
			Cluster: cluster,
			Http:    http,
			Db:      db,
		})
	}
	return connections, nil
}
