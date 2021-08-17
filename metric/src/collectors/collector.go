package collectors

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

type CCollector struct {
	param      *commands.ExecutorParameter
	connectors []connection.Connector
	internal   []*Builder
	metrics    []*Builder
}

func (c *CCollector) AddInternal(b []*Builder) *CCollector {
	c.internal = append(c.internal, b...)
	return c
}

func (c *CCollector) AddMetrics(b []*Builder) *CCollector {
	c.metrics = append(c.metrics, b...)
	return c
}

func (c *CCollector) Describe(channel chan<- *prometheus.Desc) {
	for _, metric := range c.metrics {
		if metric.Desc != nil {
			channel <- metric.Desc
		}
	}
}

func (c *CCollector) Collect(channel chan<- prometheus.Metric) {
	for _, builder := range c.internal {
		for _, metric := range builder.Builder(builder.Desc, nil, c.param) {
			channel <- metric
		}
	}

	for _, connector := range c.connectors {
		for _, builder := range c.metrics {
			for _, metric := range builder.Builder(builder.Desc, connector, c.param) {
				channel <- metric
			}
		}
	}
}

func New(param *commands.ExecutorParameter, connectors []connection.Connector) *CCollector {
	return &CCollector{
		param:      param,
		connectors: connectors,
		internal:   make([]*Builder, 0),
		metrics:    make([]*Builder, 0),
	}
}
