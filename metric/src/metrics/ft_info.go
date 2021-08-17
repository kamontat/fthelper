package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var FTInfo = collectors.NewMetrics(
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "build", "info"),
		"Information relate to freqtrade, 0 meaning server is down",
		FreqtradeLabel(),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var up float64 = 0
		if freqtrade.ToPing(connector) {
			up = 1
		}

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			up,
			FreqtradeLabelValues(connector)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "", "state"),
		"Freqtrade run state",
		append(FreqtradeLabel(), "text"),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var status, _ = freqtrade.ToStatus(connector)

		var labels = append(FreqtradeLabelValues(connector), status.StateStr())
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			float64(status.StateInt()),
			labels...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "", "mode"),
		"Freqtrade mode",
		append(FreqtradeLabel(), "text"),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var status, _ = freqtrade.ToStatus(connector)

		var labels = append(FreqtradeLabelValues(connector), status.RunMode)
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			float64(status.ModeInt()),
			labels...,
		)}
	}),
)
