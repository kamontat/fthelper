package freqtrade

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/maps"
)

const VERSION_CONST = "version"

func NewVersion() *version {
	return &version{}
}

type version struct{}

func (v *version) Name() string {
	return VERSION_CONST
}

func (p *version) Build(connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	var target = make(maps.Mapper)
	err := connection.Http.GET(p.Name(), &target)
	if err != nil {
		return nil, err
	}

	return target.Se("version")
}

func ToVersion(connector connection.Connector) string {
	raw, err := connector.Connect(VERSION_CONST)
	if err != nil {
		return "v0.0.0"
	}
	return raw.(string)
}
