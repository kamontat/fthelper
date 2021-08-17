package freqtrade

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
)

const PING_CONST = "ping"

func NewPing() *ping {
	return &ping{}
}

type ping struct{}

func (p *ping) Name() string {
	return PING_CONST
}

func (p *ping) Build(connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	var target = make(map[string]interface{})
	err := connection.Http.GET(p.Name(), &target)
	return target["status"] == "pong", err
}

func ToPing(connector connection.Connector) bool {
	raw, err := connector.Connect(PING_CONST)
	if err != nil {
		return false
	}
	return raw.(bool)
}
