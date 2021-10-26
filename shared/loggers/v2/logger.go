package loggers

import (
	"fmt"
	"strings"
	"time"
)

// Logger is printable object where it has level and format of the output
// Normally, you going to use this kind of output for enduser to understand what happen
type Logger struct {
	Name    string
	printer *Printer
}

func (l *Logger) IsPrintable(lvl Level) bool {
	return level >= lvl
}

func (l *Logger) format(lvl, format string, msg ...interface{}) string {
	// format syntax datetime: `02-01-2006 15:04:05`
	var datetime = time.Now().Format("15:04:05")

	var arr = make([]interface{}, 3)
	arr[0] = datetime
	arr[1] = l.Name
	arr[2] = strings.ToUpper(lvl)
	arr = append(arr, msg...)

	return fmt.Sprintf("%s %-20s [%-5s] | "+format, arr...)
}

func (l *Logger) Debug(format string, msg ...interface{}) {
	if l.IsPrintable(DEBUG) {
		l.printer.Print(l.format("debug", format, msg...))
	}
}

func (l *Logger) Info(format string, msg ...interface{}) {
	if l.IsPrintable(INFO) {
		l.printer.Print(l.format("info", format, msg...))
	}
}

func (l *Logger) Warn(format string, msg ...interface{}) {
	if l.IsPrintable(WARN) {
		l.printer.Print(l.format("warn", format, msg...))
	}
}

func (l *Logger) ErrorString(format string, msg ...interface{}) {
	if l.IsPrintable(ERROR) {
		l.printer.Print(l.format("error", format, msg...))
	}
}

func (l *Logger) Error(err error) {
	if err != nil {
		l.ErrorString(err.Error())
	}
}

func (l *Logger) Log(format string, msg ...interface{}) {
	if level != SILENT {
		l.printer.Print(fmt.Sprintf(format, msg...))
	}
}

func (l *Logger) Line() {
	l.Log(LINE)
}

func (l *Logger) NewLine() {
	l.Log(EMPTY)
}

func New(name string, printer *Printer) *Logger {
	return &Logger{
		Name:    name,
		printer: printer,
	}
}
