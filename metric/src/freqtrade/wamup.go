package freqtrade

import "time"

func Warmup(conn *Connection) time.Duration {
	var start = time.Now()

	NewStatus(conn)
	NewVersion(conn)
	NewCount(conn)
	NewBalance(conn)
	NewProfit(conn)
	NewLocks(conn)
	NewWhitelist(conn)
	NewPerformance(conn)
	NewLogs(conn)

	return time.Since(start)
}
