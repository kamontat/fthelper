package loggers

import (
	"math"
	"strings"
	"text/tabwriter"
)

type Table struct {
	size   uint
	logger *Logger

	writer *tabwriter.Writer
}

func (l *Table) Init() *Table {
	var lineSize = len(LINE)
	var size = int(float64(lineSize) / float64(l.size))
	var min = math.Min(float64(size), float64(4))

	l.writer = tabwriter.NewWriter(l.logger.writer, int(min), size, 2, ' ', 0)
	return l
}

func (l *Table) ToMsg(msg ...string) string {
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

func (l *Table) Header(msg ...string) {
	l.logger.FLog(l.writer, l.ToMsg(msg...))
}

func (l *Table) Row(msg ...string) {
	l.logger.FLog(l.writer, l.ToMsg(msg...))
}

func (l *Table) End() error {
	return l.writer.Flush()
}
