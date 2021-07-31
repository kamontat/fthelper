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
	NewCount(conn)
	NewLocks(conn)
	NewLogs(conn)
	NewPerformance(conn)
	NewProfit(conn)
	NewStatus(conn)
	NewVersion(conn)
	NewWhitelist(conn)
	NewStat(conn)

	err.And(warmupDailyPerformance(start, conn, logger))

	if err.HasError() {
		logger.Warn("Get %d warmup error", err.Length())
		logger.Warn(err.String())
	}

	return time.Since(start), err
}
