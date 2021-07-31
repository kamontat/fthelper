package freqtrade

import (
	"time"

	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/loggers"
)

const (
	CACHE_DAILY_PERFORMANCE_BALANCE = "daily-performance-balance"
)

func warmupDailyPerformance(start time.Time, conn *Connection, logger *loggers.Logger) error {
	var midnight = 0
	var expireAt = 24 * time.Hour

	var data *caches.Data = nil
	if caches.Global.Has(CACHE_DAILY_PERFORMANCE_BALANCE) {
		data = caches.Global.Get(CACHE_DAILY_PERFORMANCE_BALANCE)
	} else if start.Hour() == midnight {
		logger.Info("initial daily performance balance")

		data = caches.NewData(CACHE_DAILY_PERFORMANCE_BALANCE, func(o interface{}) (interface{}, error) {
			return FetchBalance(conn)
		}, expireAt)

		var err = caches.Global.SetData(data) // add data to cache service
		if err != nil {
			return err
		}
	}

	if data != nil {
		var _, err = data.Fetch()
		if err != nil {
			return err
		}
	}

	return nil
}
