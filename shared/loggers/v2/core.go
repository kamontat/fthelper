package loggers

import "github.com/kamontat/fthelper/shared/utils"

var level Level = INFO
var printer *Printer = NewDefaultPrinter()
var table *Table = NewDefaultTable()

var storage = make(map[string]*Logger)

func SetLevel(l float64) {
	level = ToLevel(l)
}

func GetLevel() Level {
	return level
}

func GetPrinter() *Printer {
	return printer
}

func GetTable() *Table {
	return table
}

func Get(names ...string) *Logger {
	name := utils.JoinString(":", names...)
	if storage[name] == nil {
		storage[name] = New(name, printer)
	}

	return storage[name]
}
