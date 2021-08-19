package metrics

import (
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
			FreqtradeLabel(),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var logs, _ = freqtrade.ToLogs(connector)
			var labels = FreqtradeLabelValues(connector)

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
			FreqtradeLabel(),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var logs, _ = freqtrade.ToLogs(connector)
			var labels = FreqtradeLabelValues(connector)

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
			append(FreqtradeLabel(), "level"),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var logs, _ = freqtrade.ToLogLevel(connector)
			var metrics = make([]prometheus.Metric, 0)
			for key, value := range logs {
				var labels = append(FreqtradeLabelValues(connector), key)
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
