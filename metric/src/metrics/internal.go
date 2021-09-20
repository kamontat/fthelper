package metrics

import (
	"github.com/kamontat/fthelper/metric/v4/src/collectors"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/constants"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/kamontat/fthelper/shared/utils"
	"github.com/prometheus/client_golang/prometheus"
)

var Internal = collectors.NewMetrics(
	collectors.NewMetric(prometheus.NewDesc(
		prometheus.BuildFQName("fthelper", "build", "info"),
		"fthelper information. Value will always change when new version is deployed.",
		[]string{"version", "commit", "date"},
		nil,
	), func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
		var number = utils.VersionNumber(param.Meta.Version)
		return []prometheus.Metric{prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			number,
			param.Meta.Version,
			param.Meta.Commit,
			param.Meta.Date,
		)}
	}),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("fthelper", "internal", "cache_size"),
			"Total keys we store on cache service, including expired ones",
			[]string{"type"},
			nil,
		),
		func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.CounterValue,
				float64(param.Cache.Size()),
				"local",
			), prometheus.MustNewConstMetric(
				desc,
				prometheus.CounterValue,
				float64(caches.Global.Size()),
				"global",
			)}
		},
	),
	collectors.NewMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName("fthelper", "internal", "warmup"),
			"Success rate of warmup process (0.0-1.0). Negative number meaning warmup is disabled",
			nil,
			nil,
		),
		func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var data = caches.Global.Get(constants.WARMUP_SUCCEESS_RATE).Data
			if data == nil {
				return emptyMetrics
			}

			return []prometheus.Metric{prometheus.MustNewConstMetric(
				desc,
				prometheus.GaugeValue,
				data.(float64),
			)}
		},
	),
	collectors.NewRawMetric(
		func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric {
			var cache = caches.Global
			var data = cache.Get(constants.WARMUP_DURATIONS)
			if !data.IsExist() {
				return []prometheus.Metric{}
			}

			var durations = data.Data.([]interface{})
			var histogram = prometheus.NewHistogram(prometheus.HistogramOpts{
				Namespace: "fthelper",
				Subsystem: "internal",
				Name:      "warmup_duration",
				Help:      "Warmup duration in milliseconds. (kept only last 1000)",
				Buckets:   prometheus.ExponentialBuckets(20, 2, 8),
			})

			for _, ms := range durations {
				histogram.Observe(float64(ms.(int64)))
			}

			return []prometheus.Metric{histogram}
		},
	),
)
