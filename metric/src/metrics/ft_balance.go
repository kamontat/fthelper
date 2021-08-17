package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

var FTBalance = collectors.NewMetrics(
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("freqtrade", "crypto", "balance"),
		"Current crypto balance in exchange",
		append(FreqtradeLabel(), "stake"),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var balance, err = freqtrade.ToBalance(connector)
		if err != nil {
			param.Logger.Warn("freqtrade.balance return empty: %v", err)
			return emptyMetrics
		}

		var labels = append(FreqtradeLabelValues(connector), balance.CryptoSymbol)
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
		append(FreqtradeLabel(), "stake"),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var balance, err = freqtrade.ToBalance(connector)
		if err != nil {
			param.Logger.Warn("freqtrade.balance return empty: %v", err)
			return emptyMetrics
		}

		var labels = append(FreqtradeLabelValues(connector), balance.FiatSymbol)
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
		append(FreqtradeLabel(), "currency", "stake"),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var balance, err = freqtrade.ToBalance(connector)
		if err != nil {
			param.Logger.Warn("freqtrade.balance return empty: %v", err)
			return emptyMetrics
		}

		var metrics = make([]prometheus.Metric, 0)
		for _, currency := range balance.Currencies {
			var labels = append(FreqtradeLabelValues(connector), currency.Symbol, currency.StakeSymbol)
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
		append(FreqtradeLabel(), "from", "to"),
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var balance, err = freqtrade.ToBalance(connector)
		if err != nil {
			param.Logger.Warn("freqtrade.balance return empty: %v", err)
			return emptyMetrics
		}

		var rate = balance.FiatValue / balance.CryptoValue
		var labels = append(FreqtradeLabelValues(connector), balance.CryptoSymbol, balance.FiatSymbol)
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			rate,
			labels...,
		)}
	}),
)
