package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var FTPerformance = collectors.NewMetrics(
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("freqtrade", "perf", "daily"),
			"Profit calculate by balance from yesterday and today (update once a day).",
			FreqtradeLabel(),
			nil,
		), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			// TODO: correct how daily performance calculate
			// var data = caches.Global.Get(freqtrade.CACHE_DAILY_PERFORMANCE_BALANCE)

			// var balance, _ = freqtrade.ToBalance(connector)
			// var previous = freqtrade.NewBalance()
			// if data.IsExist() {
			// 	previous = data.Data.(*freqtrade.Balance)
			// }

			var labels = FreqtradeLabelValues(connector)
			// var value, ok = aggregators.PercentChange(previous.CryptoValue, balance.CryptoValue)
			var value, ok = float64(0), true
			if !ok {
				param.Logger.Info("skip 'perf_daily' because previous is not exist")
				return emptyMetrics
			}

			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				value,
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
