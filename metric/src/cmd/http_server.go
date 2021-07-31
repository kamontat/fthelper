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

func HttpServer(p *commands.ExecutorParameter, conn *freqtrade.Connection) error {
	var serverPort = p.Config.Mi("server").No("port", 8090)
	var serverPath = p.Config.Mi("server").So("metric-path", "/metrics")

	prometheus.MustRegister(
		collectors.New(p, conn, metrics.FT),
		collectors.New(p, conn, metrics.FTBalance),
		collectors.New(p, conn, metrics.FTTrade),
		collectors.New(p, conn, metrics.FTPair),
		collectors.New(p, conn, metrics.FTPerformance),
		collectors.New(p, conn, metrics.FTLock),
		collectors.New(p, conn, metrics.FTLog),
		collectors.New(p, conn, metrics.Info),
		collectors.New(p, conn, metrics.FTInternal),
		collectors.New(p, conn, metrics.Internal),
	)
	http.Handle(serverPath, promhttp.Handler())
	http.HandleFunc("/version", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "%s: %s (%s)", p.Meta.Name, p.Meta.Version, p.Meta.Commit)
	})

	p.Logger.Info("Start server at 0.0.0.0:%.0f", serverPort)
	return http.ListenAndServe(fmt.Sprintf(":%.0f", serverPort), nil)
}
