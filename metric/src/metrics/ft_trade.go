package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var FTTrade = collectors.NewMetrics(
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "open_stake"),
		"Total stake in crypto for all opened trade.",
		FreqtradeLabel(),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {

		var count, _ = freqtrade.ToCount(connector)

		var labels = FreqtradeLabelValues(connector)
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			count.TotalStake,
			labels...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "open_current"),
		"Currently open trade",
		FreqtradeLabel(),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {

		var count, _ = freqtrade.ToCount(connector)

		var labels = FreqtradeLabelValues(connector)
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			float64(count.Current),
			labels...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "open_max"),
		"Maximum allow to open trade at the time.",
		FreqtradeLabel(),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {

		var count, _ = freqtrade.ToCount(connector)

		var labels = FreqtradeLabelValues(connector)
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			float64(count.Max),
			labels...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "total"),
		"Total trades, increasing both open and close",
		FreqtradeLabel(),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var profit, _ = freqtrade.ToProfit(connector)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.CounterValue,
			float64(profit.TotalTrades),
			FreqtradeLabelValues(connector)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "close"),
		"Total closed trades.",
		FreqtradeLabel(),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var profit, _ = freqtrade.ToProfit(connector)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.CounterValue,
			float64(profit.ClosedTrades),
			FreqtradeLabelValues(connector)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "winning"),
		"How many trade are winning as counter",
		FreqtradeLabel(),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var profit, _ = freqtrade.ToProfit(connector)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.CounterValue,
			float64(profit.WinTrades),
			FreqtradeLabelValues(connector)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "losing"),
		"How many trade are losing as counter",
		FreqtradeLabel(),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {

		var profit, _ = freqtrade.ToProfit(connector)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.CounterValue,
			float64(profit.LossTrades),
			FreqtradeLabelValues(connector)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "avg_duration_seconds"),
		"Average closing trade in seconds.",
		FreqtradeLabel(),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var profit, _ = freqtrade.ToProfit(connector)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			profit.GetAverageDuration().Seconds(),
			FreqtradeLabelValues(connector)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "first_seconds"),
		"First trade since unix epoch in seconds.",
		FreqtradeLabel(),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {

		var profit, _ = freqtrade.ToProfit(connector)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			float64(profit.FirstTradeTimestamp)/float64(1000),
			FreqtradeLabelValues(connector)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "latest_seconds"),
		"Latest trade since unix epoch in seconds.",
		FreqtradeLabel(),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {

		var profit, _ = freqtrade.ToProfit(connector)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			float64(profit.LastTradeTimestamp)/float64(1000),
			FreqtradeLabelValues(connector)...,
		)}
	}),
)
