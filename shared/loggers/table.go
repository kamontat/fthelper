package loggers

import (
	"math"
	"os"
	"strings"
	"text/tabwriter"
)

type TableLogger struct {
	size   uint
	logger *Logger

	writer *tabwriter.Writer
}

func (l *TableLogger) Init() *TableLogger {
	var lineSize = len(LINE)
	var size = int(float64(lineSize) / float64(l.size))
	var min = math.Min(float64(size), float64(4))

	l.writer = tabwriter.NewWriter(os.Stdout, int(min), size, 2, ' ', 0)
	return l
}

func (l *TableLogger) ToMsg(msg ...string) string {
	var str strings.Builder

	j := 0
	for i := uint(0); i < l.size; i++ {
		if i > 0 {
			str.WriteRune('\t')
		}

		if j >= len(msg) {
			str.WriteString("")
		} else {
			str.WriteString(msg[i])
		}
		j++
	}

	return str.String()
}

func (l *TableLogger) Header(msg ...string) {
	l.logger.FLog(l.writer, l.ToMsg(msg...))
}

func (l *TableLogger) Row(msg ...string) {
	l.logger.FLog(l.writer, l.ToMsg(msg...))
}

func (l *TableLogger) End() error {
	return l.writer.Flush()
}
