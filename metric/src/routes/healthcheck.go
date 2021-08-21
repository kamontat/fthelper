package routes

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/commandline/commands"
)

var HealthCheck = &Route{
	Path: "/healthcheck",
	Handler: func(p *commands.ExecutorParameter, connectors []connection.Connector) (int, interface{}) {
		var pass = true
		for _, connector := range connectors {
			var err = connector.Initial()
			if err != nil {
				pass = false
				p.Logger.Error("Connector ('%s') has a error: %v", connector.Cluster(), err)
			}
		}

		if pass {
			return 200, "done"
		}

		return 500, "fail"
	},
}
