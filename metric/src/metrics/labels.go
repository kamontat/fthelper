package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
)

func FreqtradeLabel() []string {
	return []string{
		"cluster",
		"strategy",
		"version",
	}
}

func FreqtradeLabelValues(connector connection.Connector) []string {
	var v = freqtrade.ToVersion(connector)
	var s, err = freqtrade.ToStatus(connector)
	var strategy = "unknown"
	if err == nil {
		strategy = s.Strategy
	}

	return []string{
		connector.Cluster(),
		strategy,
		v,
	}
}
