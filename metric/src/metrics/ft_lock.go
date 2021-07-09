package metrics

import (
	"github.com/kamontat/fthelper/metric/src/collectors"
	"github.com/kamontat/fthelper/metric/src/connection"
	"github.com/kamontat/fthelper/metric/src/freqtrade"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var FTLock = collectors.NewMetrics(
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "lock", "count"),
		"Current active lock data.",
		freqtrade.SummaryLabel(),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var locks = freqtrade.NewLocks(connection)

		var labels = freqtrade.NewSummary(connection, param.Cache)
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			float64(locks.Count),
			labels...,
		)}
	}),
)
