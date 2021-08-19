package freqtrade

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
)

const PERF_CONST = "performance"

func NewPerformance() *performances {
	return &performances{
		Data: make([]performance, 0),
	}
}

type performance struct {
	Pair      string  `json:"pair"`
	Profit    float64 `json:"profit"`
	ProfitAbs float64 `json:"profit_abs"`
	Count     int64   `json:"count"`
}

type performances struct {
	Data []performance
}

func (p *performances) Name() string {
	return PERF_CONST
}

func (p *performances) Build(connector connection.Connector, connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	err := connection.Http.GET(p.Name(), &p.Data)
	return p, err
}

func ToPerformance(connector connection.Connector) (*performances, error) {
	raw, err := connector.Connect(PERF_CONST)
	if err != nil {
		return nil, err
	}
	return raw.(*performances), nil
}
