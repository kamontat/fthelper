package routes

import (
	"fmt"
	"net/http"

	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/commandline/commands"
)

type Route struct {
	Path    string
	Handler func(p *commands.ExecutorParameter, connectors []connection.Connector) (int, interface{})
}

func Apply(p *commands.ExecutorParameter, connectors []connection.Connector, routes ...*Route) {
	for _, route := range routes {
		http.HandleFunc(route.Path, func(rw http.ResponseWriter, r *http.Request) {
			var code, response = route.Handler(p, connectors)

			rw.WriteHeader(code)
			fmt.Fprintln(rw, response)
		})
	}
}
