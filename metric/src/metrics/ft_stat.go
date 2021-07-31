package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var FT = collectors.NewMetrics(
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "stat", "sell_reason"),
		"Sell reason wins/draws/losses number.",
		append(freqtrade.SummaryLabel(), "reason", "type"),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var stats = freqtrade.NewStat(connection)

		var metrics = make([]prometheus.Metric, 0)
		for name, stat := range stats.Reasons {
			metrics = append(metrics, prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				float64(stat.Win),
				append(freqtrade.NewSummary(connection, param.Cache), name, "win")...,
			), prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				float64(stat.Draw),
				append(freqtrade.NewSummary(connection, param.Cache), name, "draw")...,
			), prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				float64(stat.Loss),
				append(freqtrade.NewSummary(connection, param.Cache), name, "loss")...,
			))
		}
		return metrics
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "stat", "duration"),
		"Current average wins/draws/losses duration",
		append(freqtrade.SummaryLabel(), "type"),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var stats = freqtrade.NewStat(connection)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			stats.WinDuration(),
			append(freqtrade.NewSummary(connection, param.Cache), "win")...,
		), prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			stats.DrawDuration(),
			append(freqtrade.NewSummary(connection, param.Cache), "draw")...,
		), prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			stats.LossDuration(),
			append(freqtrade.NewSummary(connection, param.Cache), "loss")...,
		)}
	}),
)
