package freqtrade

import "github.com/kamontat/fthelper/shared/caches"

type Summary struct {
	Version string
}

func SummaryLabel() []string {
	return []string{
		"cluster",
		"strategy",
		"version",
	}
}

func NewSummary(conn *Connection, cache *caches.Service) []string {
	var v = NewVersion(conn)
	var s = NewStatus(conn)

	return []string{
		conn.Cluster,
		s.Strategy,
		v,
	}
}
