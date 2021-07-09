package loggers

import (
	"github.com/kamontat/fthelper/shared/utils"
)

var level LoggerLevel = INFO
var storage = make(map[string]*Logger)

func new(name string) *Logger {
	return &Logger{
		Name: name,
	}
}

func IsDebug() bool {
	return level == DEBUG
}

func Level(l LoggerLevel) {
	level = l
}

func Get(names ...string) *Logger {
	name := utils.JoinString(":", names...)
	if storage[name] == nil {
		storage[name] = new(name)
	}

	return storage[name]
}
