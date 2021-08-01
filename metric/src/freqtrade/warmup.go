package freqtrade

import (
	"time"

	"github.com/kamontat/fthelper/shared/errors"
	"github.com/kamontat/fthelper/shared/loggers"
)

func Warmup(conn *Connection) (time.Duration, *errors.Handler) {
	var logger = loggers.Get("freqtrade", "warmup")

	var err = errors.New()
	var start = time.Now()

	err.AndD(FetchBalance(conn))
	err.AndD(FetchCount(conn))
	err.AndD(FetchCount(conn))
	err.AndD(FetchLocks(conn))
	err.AndD(FetchLogs(conn))
	err.AndD(FetchPerformance(conn))
	err.AndD(FetchProfit(conn))
	err.AndD(FetchStat(conn))
	err.AndD(FetchStatus(conn))
	err.AndD(FetchVersion(conn))
	err.AndD(FetchWhitelist(conn))

	err.And(warmupDailyPerformance(start, conn, logger))
	err.And(warmupCurrencyRate(conn, logger))

	if err.HasError() {
		logger.Warn("Get %d warmup error", err.Length())
		logger.Warn(err.String())
	}

	return time.Since(start), err
}
