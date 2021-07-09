package loggers

type LoggerLevel int8

const (
	SILENT LoggerLevel = iota
	ERROR
	WARN
	INFO
	DEBUG
)
