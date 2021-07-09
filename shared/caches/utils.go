package caches

import "time"

// parseDuration will panic if format error
func parseDuration(s string) time.Duration {
	if s == "" {
		return -1
	}

	d, e := time.ParseDuration(s)
	if e != nil {
		panic(e)
	}

	return d
}
