module github.com/kamontat/fthelper/metric

go 1.16

replace github.com/kamontat/fthelper/shared v0.0.0 => ../shared

require (
	github.com/kamontat/fthelper/shared v0.0.0
	github.com/prometheus/client_golang v1.11.0
)
