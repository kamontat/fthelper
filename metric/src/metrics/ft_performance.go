package metrics

import (
	"fmt"

	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var FTPerformance = collectors.NewMetrics(
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "perf", "minute"),
			"Profit calculate by balance from last minute to now (update every minute).",
			FreqtradeLabel(),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var perf, err = freqtrade.ToSchedulerPerformance(connector)
			if err != nil {
				fmt.Println(err)
			}

			var labels = FreqtradeLabelValues(connector)

			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				perf.Minute,
				labels...,
			)}
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "perf", "hourly"),
			"Profit calculate by balance from last hour to now (update once every hour).",
			FreqtradeLabel(),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var perf, _ = freqtrade.ToSchedulerPerformance(connector)
			var labels = FreqtradeLabelValues(connector)

			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				perf.Hourly,
				labels...,
			)}
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "perf", "daily"),
			"Profit calculate by balance from yesterday to now (update once everyday).",
			FreqtradeLabel(),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var perf, _ = freqtrade.ToSchedulerPerformance(connector)
			var labels = FreqtradeLabelValues(connector)

			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				perf.Daily,
				labels...,
			)}
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "perf", "monthly"),
			"Profit calculate by balance from last month to now (update once every month).",
			FreqtradeLabel(),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var perf, _ = freqtrade.ToSchedulerPerformance(connector)
			var labels = FreqtradeLabelValues(connector)

			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				perf.Monthly,
				labels...,
			)}
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "perf", "realized"),
			"Realized profit amount (included only closed trades).",
			append(FreqtradeLabel(), "stake"),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var balance, err1 = freqtrade.ToBalance(connector)
			var profit, err2 = freqtrade.ToProfit(connector)

			// handle when fetching return error
			if err1 != nil || err2 != nil {
				return emptyMetrics
			}

			var labels = FreqtradeLabelValues(connector)
			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				profit.RealizedCryptoProfit,
				append(labels, balance.CryptoSymbol)...,
			), prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				profit.RealizedFiatProfit,
				append(labels, balance.FiatSymbol)...,
			)}
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "perf", "realized_pct"),
			"Realized profit percentage (0-1).",
			FreqtradeLabel(),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var profit, _ = freqtrade.ToProfit(connector)
			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				profit.RealizedPercentProfit,
				FreqtradeLabelValues(connector)...,
			)}
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "perf", "unrealized"),
			"Unrealized profit amount (included both opened/closed trades).",
			append(FreqtradeLabel(), "stake"),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var balance, err1 = freqtrade.ToBalance(connector)
			var profit, err2 = freqtrade.ToProfit(connector)

			// handle when fetching return error
			if err1 != nil || err2 != nil {
				return emptyMetrics
			}

			var labels = FreqtradeLabelValues(connector)
			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				profit.UnrealizedCryptoProfit,
				append(labels, balance.CryptoSymbol)...,
			), prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				profit.UnrealizedFiatProfit,
				append(labels, balance.FiatSymbol)...,
			)}
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "perf", "unrealized_pct"),
			"Unrealized profit percentage (0-1).",
			FreqtradeLabel(),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var profit, _ = freqtrade.ToProfit(connector)
			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				profit.RealizedPercentProfit,
				FreqtradeLabelValues(connector)...,
			)}
		},
	),
)
