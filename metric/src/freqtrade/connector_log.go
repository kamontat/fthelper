package freqtrade

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

const (
	LOG_ERROR = "ERROR"
	LOG_WARN  = "WARNING"
	LOG_INFO  = "INFO"
)

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

func EmptyLogs() *Logs {
	return &Logs{
		Total: 0,
		Valid: 0,
		List:  make([]*Log, 0),
	}
}

func buildLogs(mapper maps.Mapper, log *loggers.Logger) *Logs {
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

			if logMessage.Level == LOG_ERROR {
				log.Warn(fmt.Sprintf("Found error from freqtrade: %v", logMessage.Message))
			}
			result = append(result, logMessage)
		} else {
			log.Warn(fmt.Sprintf("found log message that cannot pass (%v)", raw))
		}
	}

	return &Logs{
		Total: int(total),
		Valid: len(result),
		List:  result,
	}
}

func NewLogs(conn *Connection) *Logs {
	var name, logger = ConnectorLog(API_LOG)
	var _, expireAt, query = Connector(conn, name)

	if data, err := conn.Cache(name, expireAt, func() (interface{}, error) {
		var target = make(maps.Mapper)
		err := conn.GET(name, query, &target)
		return buildLogs(target, logger), err
	}); err == nil && data != nil {
		return data.(*Logs)
	}

	return EmptyLogs()
}
