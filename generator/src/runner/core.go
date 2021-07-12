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

		// 2. 'withCluster=true' will run multiple execution with single cluster; this will add cluster as string
		//    'withCluster=false' will run single execution with multiple cluster; this will add clusters as []string
		var raws = config.Mi("internal").Ai("clusters")
		if len(raws) < 1 {
			raws = append(raws, "") // add default empty cluster
		}

		var name = data.Si("type")
		var defaultWithCluster = true
		if name == "copy" {
			defaultWithCluster = false
		}

		if data.Bo("withCluster", defaultWithCluster) {
			for _, raw := range raws {
				var cluster = raw.(string)

				// 1. add cluster to .data
				var newData = data.Copy().Set("cluster", cluster)

				// 2. add data to config (.data)
				var merger = maps.Merger(config).Add(maps.New().Set("data", newData))

				log.Debug("override data with cluster: %s", cluster) // this might be remove because cluster never changes
				merger.Add(config.Mi("_").Mi(cluster))

				var err = executor(&ExecutorParameter{
					Name:     name,
					Data:     newData,
					Config:   merger.Merge(),
					FsConfig: config.Mi("fs"),
					Logger:   log,
				})

				if err != nil {
					return err
				}
			}
			// finished
			return nil
		}

		var clusters = make([]string, 0)
		for _, raw := range raws {
			clusters = append(clusters, raw.(string))
		}

		// 1. add clusters to .data
		var newData = data.Copy().Set("clusters", clusters)
		var merger = maps.Merger(config).Add(maps.New().Set("data", newData))

		return executor(&ExecutorParameter{
			Name:     name,
			Data:     newData,
			Config:   merger.Merge(),
			FsConfig: config.Mi("fs"),
			Logger:   log,
		})
	}).Input(data)
}
