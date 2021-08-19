package freqtrade

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
)

const WHITELIST_CONST = "whitelist"

func NewWhitelist() *whitelist {
	return &whitelist{
		Length: 0,
	}
}

type whitelist struct {
	Length int64
}

func (w *whitelist) Name() string {
	return WHITELIST_CONST
}

func (w *whitelist) Build(connector connection.Connector, connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	err := connection.Http.GET(w.Name(), w)
	return w, err
}

func ToWhitelist(connector connection.Connector) (*whitelist, error) {
	raw, err := connector.Connect(WHITELIST_CONST)
	if err != nil {
		return nil, err
	}
	return raw.(*whitelist), nil
}
