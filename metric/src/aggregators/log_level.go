package aggregators

import "github.com/kamontat/fthelper/metric/v4/src/freqtrade"

// LogLevel will aggregator log data by level and return map of log level
func LogLevel(logs *freqtrade.Logs) map[string]float64 {
	var mapper = make(map[string]float64)
	for _, log := range logs.List {
		mapper[log.Level] += 1
	}

	return mapper
}
