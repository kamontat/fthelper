package cmd

import (
	"fmt"
	"net/http"

	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/metric/v4/src/metrics"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HttpServer(parameters *commands.ExecutorParameter, conn *freqtrade.Connection) error {
	var serverPort = parameters.Config.Mi("server").Io("port", 8090)
	var serverPath = parameters.Config.Mi("server").So("metric-path", "/metrics")

	prometheus.MustRegister(
		collectors.New(parameters, conn, metrics.FT),
		collectors.New(parameters, conn, metrics.FTTrade),
		collectors.New(parameters, conn, metrics.FTPair),
		collectors.New(parameters, conn, metrics.FTLock),
		collectors.New(parameters, conn, metrics.FTLog),
		collectors.New(parameters, conn, metrics.Info),
		collectors.New(parameters, conn, metrics.FTInternal),
		collectors.New(parameters, conn, metrics.Internal),
	)
	http.Handle(serverPath, promhttp.Handler())
	http.HandleFunc("/version", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "%s: %s (%s)", parameters.Meta.Name, parameters.Meta.Version, parameters.Meta.Commit)
	})

	parameters.Logger.Info("Start server at 0.0.0.0:%d", serverPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
}
