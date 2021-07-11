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
		var data = i.Input().(maps.Mapper)

		// 1. add .data
		var merger = maps.Merger(config).Add(maps.New().Set("data", data))

		// 2. add override cluster if 'withCluster' is true
		if data.Bo("withCluster", true) {
			var cluster = config.Mi("internal").Si("cluster")
			log.Debug("override data with cluster: %s", cluster) // this might be remove because cluster never changes
			merger.Add(config.Mi("_").Mi(cluster))
		}

		return executor(&ExecutorParameter{
			Name:     name,
			Data:     data,
			Config:   merger.Merge(),
			FsConfig: config.Mi("fs"),
			Logger:   log,
		})
	}).Input(data)
}
