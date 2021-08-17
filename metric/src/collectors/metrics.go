package collectors

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

type BuilderFn func(desc *prometheus.Desc, connector connection.Connector, param *commands.ExecutorParameter) []prometheus.Metric

type Builder struct {
	Desc    *prometheus.Desc
	Builder BuilderFn
}

func NewMetric(desc *prometheus.Desc, builder BuilderFn) *Builder {
	return &Builder{
		Desc:    desc,
		Builder: builder,
	}
}

func NewRawMetric(builder BuilderFn) *Builder {
	return &Builder{
		Desc:    nil,
		Builder: builder,
	}
}

func NewMetrics(metrics ...*Builder) []*Builder {
	return metrics
}
