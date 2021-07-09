package collectors

import (
	"github.com/kamontat/fthelper/metric/src/connection"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/prometheus/client_golang/prometheus"
)

type CCollector struct {
	param   *commands.ExecutorParameter
	conn    connection.Http
	metrics []*Builder
}

func (c *CCollector) Describe(channel chan<- *prometheus.Desc) {
	for _, metric := range c.metrics {
		if metric.Desc != nil {
			channel <- metric.Desc
		}
	}
}

func (c *CCollector) Collect(channel chan<- prometheus.Metric) {
	for _, builder := range c.metrics {
		for _, metric := range builder.Builder(builder.Desc, c.conn, c.param) {
			channel <- metric
		}
	}
}

func New(param *commands.ExecutorParameter, conn connection.Http, metrics []*Builder) *CCollector {
	return &CCollector{
		param:   param,
		conn:    conn,
		metrics: metrics,
	}
}
