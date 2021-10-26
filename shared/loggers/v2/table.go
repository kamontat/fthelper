package loggers

import (
	"fmt"
	"math"
	"strings"
	"text/tabwriter"
)

var (
	errTableNotInitial = fmt.Errorf("you never initial table")
)

type Table struct {
	size    uint
	printer *Printer
	writer  *tabwriter.Writer
}

func (t *Table) SetSize(size uint) *Table {
	t.size = size
	return t
}

func (t *Table) Init() *Table {
	var lineSize = len(LINE)
	var size = int(float64(lineSize) / float64(t.size))
	var min = math.Min(float64(size), float64(4))

	t.writer = tabwriter.NewWriter(t.printer.writer, int(min), size, 2, ' ', 0)
	return t
}

func (t *Table) IsInitial() bool {
	return t.writer != nil
}

func (t *Table) ToMsg(msg ...string) string {
	var str strings.Builder
	for i := 0; i < int(math.Min(float64(len(msg)), float64(t.size))); i++ {
		if i > 0 {
			str.WriteRune('\t')
		}

		str.WriteString(msg[i])
	}

	return str.String()
}

func (t *Table) Row(msg ...string) error {
	if t.IsInitial() {
		t.printer.Write(t.writer, t.ToMsg(msg...))
		return nil
	}
	return errTableNotInitial
}

func (t *Table) End() error {
	if t.IsInitial() {
		return t.writer.Flush()
	}
	return nil
}

func NewTable(printer *Printer) *Table {
	return &Table{
		size:    0,
		printer: printer,
	}
}

func NewDefaultTable() *Table {
	return NewTable(printer)
}
