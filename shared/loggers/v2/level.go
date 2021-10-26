package loggers

type Level int8

const (
	SILENT Level = iota
	ERROR
	WARN
	INFO
	DEBUG
)

func ToLevel(level float64) Level {
	if level < float64(SILENT) {
		return SILENT
	} else if level > float64(DEBUG) {
		return DEBUG
	} else {
		return Level(level)
	}
}
