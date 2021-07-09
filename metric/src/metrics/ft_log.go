package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/aggregators"
	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var FTLog = collectors.NewMetrics(
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "log", "level_total"),
			"Total log messages that we request.",
			freqtrade.SummaryLabel(),
			nil,
		), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
			var connection = freqtrade.ToConnection(conn)
			var logs = freqtrade.NewLogs(connection)

			var labels = freqtrade.NewSummary(connection, param.Cache)
			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.CounterValue,
				float64(logs.Total),
				labels...,
			)}
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "log", "level_count"),
			"Count all valid log that ftmetric can pass.",
			freqtrade.SummaryLabel(),
			nil,
		), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
			var connection = freqtrade.ToConnection(conn)
			var logs = freqtrade.NewLogs(connection)

			var labels = freqtrade.NewSummary(connection, param.Cache)
			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.CounterValue,
				float64(logs.Valid),
				labels...,
			)}
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "log", "level"),
			"How many log occurred group by level in specify time.",
			append(freqtrade.SummaryLabel(), "level"),
			nil,
		), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
			var connection = freqtrade.ToConnection(conn)
			var logs = freqtrade.NewLogs(connection)
			var aggregated = aggregators.LogLevel(logs)

			var metrics = make([]prometheus.Metric, 0)
			for key, value := range aggregated {

				var labels = append(freqtrade.NewSummary(connection, param.Cache), key)
				metrics = append(metrics, prometheus.MustNewConstMetric(
					desc,
					prometheus.CounterValue,
					value,
					labels...,
				))
			}

			return metrics
		},
	),
)
