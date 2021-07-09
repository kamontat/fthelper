package metrics

import (
	"time"

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
		freqtrade.SummaryLabel(),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var count = freqtrade.NewCount(connection)

		var labels = freqtrade.NewSummary(connection, param.Cache)
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
		freqtrade.SummaryLabel(),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var count = freqtrade.NewCount(connection)

		var labels = freqtrade.NewSummary(connection, param.Cache)
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
		freqtrade.SummaryLabel(),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var count = freqtrade.NewCount(connection)

		var labels = freqtrade.NewSummary(connection, param.Cache)
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
		freqtrade.SummaryLabel(),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var profit = freqtrade.NewProfit(connection)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.CounterValue,
			float64(profit.TotalTrades),
			freqtrade.NewSummary(connection, param.Cache)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "close"),
		"Total closed trades.",
		freqtrade.SummaryLabel(),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var profit = freqtrade.NewProfit(connection)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.CounterValue,
			float64(profit.ClosedTrades),
			freqtrade.NewSummary(connection, param.Cache)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "winning"),
		"How many trade are winning as counter",
		freqtrade.SummaryLabel(),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var profit = freqtrade.NewProfit(connection)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.CounterValue,
			float64(profit.WinTrades),
			freqtrade.NewSummary(connection, param.Cache)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "losing"),
		"How many trade are losing as counter",
		freqtrade.SummaryLabel(),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var profit = freqtrade.NewProfit(connection)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.CounterValue,
			float64(profit.LossTrades),
			freqtrade.NewSummary(connection, param.Cache)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "avg_duration_seconds"),
		"Average closing trade in seconds.",
		freqtrade.SummaryLabel(),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var profit = freqtrade.NewProfit(connection)

		var baseT, _ = time.Parse("15:04:05", "00:00:00")
		var avgT, _ = time.Parse("15:04:05", profit.AverageDuration)
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			avgT.Sub(baseT).Seconds(),
			freqtrade.NewSummary(connection, param.Cache)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "first_seconds"),
		"First trade since unix epoch in seconds.",
		freqtrade.SummaryLabel(),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var profit = freqtrade.NewProfit(connection)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			float64(profit.FirstTradeTimestamp)/float64(1000),
			freqtrade.NewSummary(connection, param.Cache)...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "trade", "latest_seconds"),
		"Latest trade since unix epoch in seconds.",
		freqtrade.SummaryLabel(),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var profit = freqtrade.NewProfit(connection)

		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			float64(profit.LastTradeTimestamp)/float64(1000),
			freqtrade.NewSummary(connection, param.Cache)...,
		)}
	}),
)
