package metrics

import (
	"github.com/kamontat/fthelper/metric/src/collectors"
	"github.com/kamontat/fthelper/metric/src/connection"
	"github.com/kamontat/fthelper/metric/src/constants"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var FTInternal = collectors.NewMetrics(
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "internal", "cache_total"),
			"How many time we call cache service for freqtrade data",
			nil,
			nil,
		),
		func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
			return callerBuilder(desc, constants.FTCONN_CACHE_TOTAL)
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "internal", "cache_miss"),
			"How many time we need to call freqtrade",
			nil,
			nil,
		),
		func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
			return callerBuilder(desc, constants.FTCONN_CACHE_MISS)
		},
	),
)
