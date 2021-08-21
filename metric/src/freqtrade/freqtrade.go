package freqtrade

import (
	"fmt"

	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/errors"
)

type Freqtrade struct {
	Connection *connection.Connection

	parent  connection.Connector
	plugins map[string]Plugin
	history map[string]*datatype.Queue
}

func (f *Freqtrade) Cluster() string {
	return f.Connection.Cluster
}

func (f *Freqtrade) Plugin(plugin Plugin) *Freqtrade {
	f.plugins[plugin.Name()] = plugin
	return f
}

func (f *Freqtrade) Initial() error {
	var err1 = f.Connection.Http.Initial()
	var err2 = f.Connection.Db.Initial()
	return errors.New().And(err1, err2).Error()
}

func (f *Freqtrade) Cleanup() error {
	return f.Connection.Db.Cleanup()
}

func (f *Freqtrade) Parent() connection.Connector {
	return f.parent
}

func (f *Freqtrade) WithParent(p connection.Connector) connection.Connector {
	f.parent = p
	return f
}

func (f *Freqtrade) Save(name string, data interface{}) connection.Connector {
	if q, ok := f.history[name]; ok {
		q.Enqueue(data)
	} else {
		f.history[name] = datatype.NewLimitQueue(10).Enqueue(data)
	}
	return f
}

func (f *Freqtrade) ConnectAll() *errors.Handler {
	var err = errors.New()
	for key := range f.plugins {
		err.AndD(f.Connect(key))
	}

	return err
}

func (f *Freqtrade) Connect(name string) (interface{}, error) {
	if plugin, ok := f.plugins[name]; ok {
		if queue, ok := f.history[name]; ok {
			return plugin.Build(f, f.Connection, queue)
		}
		return plugin.Build(f, f.Connection, datatype.NewQueue())
	}
	return nil, fmt.Errorf("'%s' is not valid name / or never implement it before", name)
}

func (f *Freqtrade) String() string {
	return f.Connection.String()
}

func New(connection *connection.Connection) connection.Connector {
	var freqtrade = &Freqtrade{
		Connection: connection,

		parent:  nil,
		plugins: make(map[string]Plugin),
		history: make(map[string]*datatype.Queue),
	}

	freqtrade.
		Plugin(NewPing()).
		Plugin(NewVersion()).
		Plugin(NewBalance()).
		Plugin(NewLogs()).
		Plugin(NewStatus()).
		Plugin(NewLocks()).
		Plugin(NewStat()).
		Plugin(NewCount()).
		Plugin(NewWhitelist()).
		Plugin(NewPerformance()).
		Plugin(NewSchedulerPerformance()).
		Plugin(NewProfit())

	return freqtrade
}
