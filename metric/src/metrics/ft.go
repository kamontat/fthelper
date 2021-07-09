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
		prometheus.BuildFQName("freqtrade", "crypto", "balance"),
		"Current crypto balance in exchange",
		append(freqtrade.SummaryLabel(), "stake"),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var balance = freqtrade.NewBalance(connection)

		var labels = append(freqtrade.NewSummary(connection, param.Cache), balance.CryptoSymbol)
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			balance.CryptoValue,
			labels...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "fiat", "balance"),
		"Current fiat balance in exchange",
		append(freqtrade.SummaryLabel(), "stake"),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var balance = freqtrade.NewBalance(connection)

		var labels = append(freqtrade.NewSummary(connection, param.Cache), balance.FiatSymbol)
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			balance.FiatValue,
			labels...,
		)}
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "crypto", "balances"),
		"Current holding cryptocurrency balance. (Number will shown as stake currency)",
		append(freqtrade.SummaryLabel(), "currency", "stake"),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var balance = freqtrade.NewBalance(connection)

		var metrics = make([]prometheus.Metric, 0)
		for _, currency := range balance.Currencies {
			var labels = append(freqtrade.NewSummary(connection, param.Cache), currency.Symbol, currency.StakeSymbol)
			metrics = append(metrics, prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				currency.EstStakeBalance,
				labels...,
			))
		}

		return metrics
	}),
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "rate", "convert"),
		"Rate from crypto to fiat currency (1 crypto = X fiat).",
		append(freqtrade.SummaryLabel(), "from", "to"),
		nil,
	), func(desc *prometheus.Desc, conn connection.Http, param *commands.ExecutorParameter) []prometheus.Metric {
		var connection = freqtrade.ToConnection(conn)
		var balance = freqtrade.NewBalance(connection)

		var rate = balance.FiatValue / balance.CryptoValue
		var labels = append(freqtrade.NewSummary(connection, param.Cache), balance.CryptoSymbol, balance.FiatSymbol)
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			rate,
			labels...,
		)}
	}),
)
