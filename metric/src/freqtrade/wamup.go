package freqtrade

import "time"

func Warmup(conn *Connection) (time.Duration, error) {
	var start = time.Now()

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

	// TODO: Add error handler properly
	return time.Since(start), nil
}
