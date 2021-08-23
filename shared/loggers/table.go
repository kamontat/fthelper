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

func (l *Table) SetSize(size uint) *Table {
	l.size = size
	return l
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
	for i := 0; i < int(math.Min(float64(len(msg)), float64(l.size))); i++ {
		if i > 0 {
			str.WriteRune('\t')
		}

		str.WriteString(msg[i])
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
