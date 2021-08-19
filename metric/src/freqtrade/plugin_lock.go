package freqtrade

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
)

const LOCK_CONST = "locks"

func NewLocks() *locks {
	return &locks{
		Count: 0,
		List:  make([]lock, 0),
	}
}

type lock struct {
	Id             int
	Active         bool
	StartTimestamp int64 `json:"lock_end_timestamp"`
	EndTimestamp   int64 `json:"lock_timestamp"`
	Pair           string
	Reason         string
}

type locks struct {
	Count int    `json:"lock_count"`
	List  []lock `json:"locks"`
}

func (l *locks) Name() string {
	return LOCK_CONST
}

func (l *locks) Build(connector connection.Connector, connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	err := connection.Http.GET(l.Name(), l)
	return l, err
}

func ToLocks(connector connection.Connector) (*locks, error) {
	raw, err := connector.Connect(LOCK_CONST)
	if err != nil {
		return nil, err
	}
	return raw.(*locks), nil
}
