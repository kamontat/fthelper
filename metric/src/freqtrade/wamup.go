package freqtrade

import "time"

func Warmup(conn *Connection) (time.Duration, error) {
	var start = time.Now()

	var err = conn.cache.FetchAll()

	return time.Since(start), err
}
