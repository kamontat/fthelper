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
			FreqtradeLabel(),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var whitelist, _ = freqtrade.ToWhitelist(connector)
			var labels = FreqtradeLabelValues(connector)

			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				float64(whitelist.Length),
				labels...,
			)}
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "pair", "profit_pct"),
			"Total percent profit per pair.",
			append(FreqtradeLabel(), "pair"),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var performances, _ = freqtrade.ToPerformance(connector)

			var metrics []prometheus.Metric = make([]prometheus.Metric, 0)
			for _, perf := range performances.Data {
				var metric = prometheus.MustNewConstMetric(
					desc,
					prometheus.GaugeValue,
					perf.Profit,
					append(FreqtradeLabelValues(connector), perf.Pair)...,
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
			append(FreqtradeLabel(), "pair"),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var performances, _ = freqtrade.ToPerformance(connector)

			var metrics []prometheus.Metric = make([]prometheus.Metric, 0)
			for _, perf := range performances.Data {
				var metric = prometheus.MustNewConstMetric(
					desc,
					prometheus.GaugeValue,
					perf.ProfitAbs,
					append(FreqtradeLabelValues(connector), perf.Pair)...,
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
			append(FreqtradeLabel(), "pair"),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var performances, _ = freqtrade.ToPerformance(connector)

			var metrics []prometheus.Metric = make([]prometheus.Metric, 0)
			for _, perf := range performances.Data {
				var metric = prometheus.MustNewConstMetric(
					desc,
					prometheus.GaugeValue,
					float64(perf.Count),
					append(FreqtradeLabelValues(connector), perf.Pair)...,
				)
				metrics = append(metrics, metric)
			}
			return metrics
		},
	),
)
