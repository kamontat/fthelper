package cmd

import (
	"fmt"
	"net/http"

	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/metric/v4/src/metrics"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HttpServer(p *commands.ExecutorParameter, connections []*freqtrade.Connection) error {
	var serverPort = p.Config.Mi("server").No("port", 8090)
	var serverPath = p.Config.Mi("server").So("metric-path", "/metrics")

	var cs []connection.Http = make([]connection.Http, 0)
	for _, c := range connections {
		cs = append(cs, c)
	}

	var collector = collectors.New(p, cs)

	collector.AddInternal(metrics.Internal)
	collector.AddInternal(metrics.FTInternal)

	collector.AddMetrics(metrics.FT)
	collector.AddMetrics(metrics.FTBalance)
	collector.AddMetrics(metrics.FTTrade)
	collector.AddMetrics(metrics.FTPair)
	collector.AddMetrics(metrics.FTPerformance)
	collector.AddMetrics(metrics.FTLock)
	collector.AddMetrics(metrics.FTLog)
	collector.AddMetrics(metrics.Info)

	prometheus.MustRegister(collector)
	http.Handle(serverPath, promhttp.Handler())
	http.HandleFunc("/version", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "%s: %s (%s)", p.Meta.Name, p.Meta.Version, p.Meta.Commit)
	})

	p.Logger.Info("Start server at 0.0.0.0:%.0f", serverPort)
	return http.ListenAndServe(fmt.Sprintf(":%.0f", serverPort), nil)
}
