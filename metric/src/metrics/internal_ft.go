package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/constants"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var InternalFT = collectors.NewMetrics(
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("fthelper", "internal", "ft_call"),
			"How many time do we call freqtrade apis",
			[]string{"cluster"},
			prometheus.Labels{
				"type": "total",
			},
		),
		func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
			var connection = freqtrade.ToConnection(conn)
			return callerClusterBuilder(desc, constants.FTCONN_CALL, connection.Cluster)
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("fthelper", "internal", "ft_call"),
			"How many time do we call freqtrade apis",
			[]string{"cluster"},
			prometheus.Labels{
				"type": "success",
			},
		),
		func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
			var connection = freqtrade.ToConnection(conn)
			return callerClusterBuilder(desc, constants.FTCONN_CALL_SUCCESS, connection.Cluster)
		},
	),
)
