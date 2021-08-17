package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var FTLock = collectors.NewMetrics(
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "lock", "count"),
		"Current active lock data.",
		FreqtradeLabel(),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var locks, _ = freqtrade.ToLocks(connector)
		var labels = FreqtradeLabelValues(connector)
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			float64(locks.Count),
			labels...,
		)}
	}),
)
