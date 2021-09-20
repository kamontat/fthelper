package caches

import (
	"time"

	"github.com/kamontat/fthelper/shared/utils"
)

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

// Join multiple cache name together
func Join(names ...string) string {
	return utils.JoinString(".", names...)
}
