package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var FTPair = collectors.NewMetrics(
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "pair", "whitelist"),
			"How many pair are whitelist currently.",
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
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "pair", "profit_pct"),
			"Total percent profit per pair.",
			append(freqtrade.SummaryLabel(), "pair"),
			nil,
		), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
			var connection = freqtrade.ToConnection(conn)
			var performances = freqtrade.NewPerformance(connection)

			var metrics []prometheus.Metric = make([]prometheus.Metric, 0)
			for _, perf := range performances {
				var metric = prometheus.MustNewConstMetric(
					desc,
					prometheus.GaugeValue,
					perf.Profit,
					append(freqtrade.NewSummary(connection, param.Cache), perf.Pair)...,
				)
				metrics = append(metrics, metric)
			}
			return metrics
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "pair", "profit_abs"),
			"Total profit per pair (as crypto currency).",
			append(freqtrade.SummaryLabel(), "pair"),
			nil,
		), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
			var connection = freqtrade.ToConnection(conn)
			var performances = freqtrade.NewPerformance(connection)

			var metrics []prometheus.Metric = make([]prometheus.Metric, 0)
			for _, perf := range performances {
				var metric = prometheus.MustNewConstMetric(
					desc,
					prometheus.GaugeValue,
					perf.ProfitAbs,
					append(freqtrade.NewSummary(connection, param.Cache), perf.Pair)...,
				)
				metrics = append(metrics, metric)
			}
			return metrics
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "pair", "count"),
			"Total buy for specify pair (including both opened and closed)",
			append(freqtrade.SummaryLabel(), "pair"),
			nil,
		), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
			var connection = freqtrade.ToConnection(conn)
			var performances = freqtrade.NewPerformance(connection)

			var metrics []prometheus.Metric = make([]prometheus.Metric, 0)
			for _, perf := range performances {
				var metric = prometheus.MustNewConstMetric(
					desc,
					prometheus.GaugeValue,
					float64(perf.Count),
					append(freqtrade.NewSummary(connection, param.Cache), perf.Pair)...,
				)
				metrics = append(metrics, metric)
			}
			return metrics
		},
	),
)
