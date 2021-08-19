package freqtrade

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/maps"
)

const (
	LOG_CONST = "logs"
	LOG_ERROR = "ERROR"
	LOG_WARN  = "WARNING"
	LOG_INFO  = "INFO"
)

func NewLogs() *Logs {
	return &Logs{
		Total: 0,
		Valid: 0,
		List:  make([]*Log, 0),
	}
}

type Log struct {
	Datetime  string
	Timestamp float64
	Namespace string
	Level     string
	Message   string
}

type Logs struct {
	Total int
	Valid int
	List  []*Log
}

func (l *Logs) Name() string {
	return LOG_CONST
}

func (l *Logs) Build(connector connection.Connector, connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	var target = make(maps.Mapper)
	err := connection.Http.GET(l.Name(), &target)
	return buildLogs(target), err
}

func ToLogs(connector connection.Connector) (*Logs, error) {
	raw, err := connector.Connect(LOG_CONST)
	if err != nil {
		return nil, err
	}
	return raw.(*Logs), nil
}

func ToLogLevel(connector connection.Connector) (map[string]float64, error) {
	var mapper = make(map[string]float64)
	var logs, err = ToLogs(connector)
	if err != nil {
		return mapper, err
	}

	for _, log := range logs.List {
		mapper[log.Level] += 1
	}

	return mapper, nil
}

func buildLogs(mapper maps.Mapper) *Logs {
	var total = mapper.Io("log_count", 0)
	var raws = mapper.Ai("logs")

	var result = make([]*Log, 0)
	for _, raw := range raws {
		var arr, ok = raw.([]interface{})
		if ok && len(arr) == 5 {
			var logMessage = &Log{
				Datetime:  arr[0].(string),
				Timestamp: arr[1].(float64),
				Namespace: arr[2].(string),
				Level:     arr[3].(string),
				Message:   arr[4].(string),
			}
			result = append(result, logMessage)
		}
	}

	return &Logs{
		Total: int(total),
		Valid: len(result),
		List:  result,
	}
}
