package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/constants"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var InternalFT = collectors.NewMetrics(
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("fthelper", "internal", "ft_call"),
			"How many time do we call freqtrade apis",
			nil,
			prometheus.Labels{
				"type": "success",
			},
		),
		func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
			return callerBuilder(desc, constants.FTCONN_CALL_SUCCESS)
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("fthelper", "internal", "ft_call"),
			"How many time do we call freqtrade apis",
			nil,
			prometheus.Labels{
				"type": "failure",
			},
		),
		func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
			return callerBuilder(desc, constants.FTCONN_CALL_FAILURE)
		},
	),
)
