package runner

import (
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func New(data maps.Mapper, config maps.Mapper, executor Executor) runners.Runner {
	var name = data.So("display", "default")
	return runners.NewRunner(name, func(i *runners.SingleInfo) error {
		return nil
	}, func(i *runners.SingleInfo) error {
		var log = loggers.Get("generator", name)

		return executor(&ExecutorParameter{
			Name:     name,
			Data:     i.Input().(maps.Mapper),
			Config:   config,
			FsConfig: config.Mi("fs"),
			Logger:   log,
		})
	}).Input(data)
}
