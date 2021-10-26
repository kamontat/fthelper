package loggers

import (
	"fmt"
	"io"
	"os"
)

// Printer is printable object where it doesn't have own level but it still respect SILENT mode
// Normally, you going to use this kind of output for custom format (table, readme etc.)
type Printer struct {
	writer io.Writer
}

func (p *Printer) Write(writer io.Writer, message interface{}) {
	fmt.Fprintln(writer, message)
}

func (p *Printer) Print(message interface{}) {
	p.Write(p.writer, message)
}

func NewPrinter(writer io.Writer) *Printer {
	return &Printer{
		writer: writer,
	}
}

func NewDefaultPrinter() *Printer {
	return NewPrinter(os.Stdout)
}
