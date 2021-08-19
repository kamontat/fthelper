package freqtrade

import (
	"time"

	"github.com/kamontat/fthelper/metric/v4/src/aggregators"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/datatype"
)

const SCHEDULER_PERF_CONST = "schedule_perf"

func NewSchedulerPerformance() *schedulerPerformance {
	return &schedulerPerformance{
		minuteBalance:  nil,
		Minute:         0,
		hourlyBalance:  nil,
		Hourly:         0,
		dailyBalance:   nil,
		Daily:          0,
		monthlyBalance: nil,
		Monthly:        0,
	}
}

type schedulerPerformance struct {
	minuteBalance  *caches.Data
	Minute         float64
	hourlyBalance  *caches.Data
	Hourly         float64
	dailyBalance   *caches.Data
	Daily          float64
	monthlyBalance *caches.Data
	Monthly        float64
}

func (p *schedulerPerformance) Name() string {
	return SCHEDULER_PERF_CONST
}

func (p *schedulerPerformance) GetBalanceChanges(data *caches.Data) float64 {
	var old *balance = nil
	if data.IsExist() {
		old = data.Data.(*balance)
	}

	var new *balance = nil
	if fetch, err := data.Fetch(); fetch && err == nil {
		new = data.Data.(*balance)
	}

	if old == nil || new == nil {
		return 0
	}

	var cal, _ = aggregators.PercentChange(old.CryptoValue, new.CryptoValue)
	return cal
}

func (p *schedulerPerformance) Build(connector connection.Connector, connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	// NOTE: This balance will call separately, meaning we didn't cache any balance result
	// If this cause high load to freqtrade, we can change to connector.Parent instead to support caching connector
	if history.Empty() {
		p.minuteBalance = caches.NewData("minute", func(o interface{}) (interface{}, error) {
			return ToBalance(connector)
		}, 1*time.Minute)
		p.hourlyBalance = caches.NewData("hourly", func(o interface{}) (interface{}, error) {
			return ToBalance(connector)
		}, 1*time.Hour)
		p.dailyBalance = caches.NewData("daily", func(o interface{}) (interface{}, error) {
			return ToBalance(connector)
		}, 24*time.Hour)
		p.monthlyBalance = caches.NewData("monthly", func(o interface{}) (interface{}, error) {
			return ToBalance(connector)
		}, 720*time.Hour)

		return p, nil // return default value (zero)
	} else {
		p.Minute = p.GetBalanceChanges(p.minuteBalance)
		p.Hourly = p.GetBalanceChanges(p.hourlyBalance)
		p.Daily = p.GetBalanceChanges(p.dailyBalance)
		p.Monthly = p.GetBalanceChanges(p.monthlyBalance)
	}

	return p, nil
}

func ToSchedulerPerformance(connector connection.Connector) (*schedulerPerformance, error) {
	raw, err := connector.Connect(SCHEDULER_PERF_CONST)
	if err != nil {
		return nil, err
	}
	return raw.(*schedulerPerformance), nil
}
