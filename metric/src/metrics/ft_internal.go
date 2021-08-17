package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var FTInternal = collectors.NewMetrics(
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "internal", "cache_total"),
			"How many time we call cache service for freqtrade data",
			[]string{"cluster"},
			nil,
		),
		func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			return callerClusterBuilder(desc, connection.CACHE_TOTAL, connector.Cluster())
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "internal", "cache_miss"),
			"How many time we need to call freqtrade",
			[]string{"cluster"},
			nil,
		),
		func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			return callerClusterBuilder(desc, connection.CACHE_MISS, connector.Cluster())
		},
	),
)
