package freqtrade

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
)

const COUNT_CONST = "count"

func NewCount() *count {
	return &count{
		Current:    -1,
		Max:        -1,
		TotalStake: 0,
	}
}

type count struct {
	Current int64
	Max     int64

	TotalStake float64 `json:"total_stake"`
}

func (c *count) Name() string {
	return COUNT_CONST
}

func (c *count) Build(connector connection.Connector, connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	err := connection.Http.GET(c.Name(), c)
	return c, err
}

func ToCount(connector connection.Connector) (*count, error) {
	raw, err := connector.Connect(COUNT_CONST)
	if err != nil {
		return nil, err
	}
	return raw.(*count), nil
}
