package freqtrade

import (
	"time"

	"github.com/kamontat/fthelper/shared/loggers"
)

func Warmup(conn *Connection) (time.Duration, error) {
	var logger = loggers.Get("freqtrade", "warmup")

	var err error = nil
	var start = time.Now()
	logger.Info("initial freqtrade connection")

	NewBalance(conn)
	NewCount(conn)
	NewLocks(conn)
	NewLogs(conn)
	NewPerformance(conn)
	NewProfit(conn)
	NewStatus(conn)
	NewVersion(conn)
	NewWhitelist(conn)
	NewStat(conn)

	err = warmupDailyPerformance(start, conn, logger)

	// TODO: Add error handler properly
	return time.Since(start), err
}
