package clusters

import (
	"strconv"

	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

// NewRunner will create cluster runner
func NewRunner(data maps.Mapper, mapper maps.Mapper, executor Executor, setting *Settings) runners.Runner {
	var name = data.So("display", "default")
	return runners.NewRunner(name, func(i *runners.SingleInfo) error {
		var config = mapper.Copy()
		var data = i.Input().(maps.Mapper).Copy()

		var index = GetIndex()
		var generatorType = data.Si("type")
		var log = loggers.Get("generator", strconv.FormatInt(int64(index), 10))

		// add .clusters to data
		var raws = config.Ai("clusters")
		if len(raws) < 1 {
			raws = append(raws, "") // add default empty cluster
		}

		// update new input value
		i.In(&ExecutorParameter{
			Index:    index,
			Name:     i.Name(),
			Type:     generatorType,
			Config:   config,
			FsConfig: config.Mi("fs"),
			Logger:   log,
			Data:     data.Set("clusters", raws),
		})

		return nil
	}, func(i *runners.SingleInfo) error {
		var param = i.Input().(*ExecutorParameter)

		param.Logger.Info("Start generate %s (%s) with data: %v", param.Name, param.Type, param.Data)
		if param.Data.Bo("withCluster", setting.DefaultWithCluster) {
			for _, raw := range param.Data.Ai("clusters") {
				var cluster = raw.(string)
				param.Logger.Debug("override data with cluster: %s", cluster) // this might be remove because cluster never changes

				// 1. Add cluster to data
				param.Data.Set("cluster", cluster)
				// 2. Add data to config
				param.Config.Set("data", param.Data)
				// 3. Override config with cluster
				var merger = maps.Merger(param.Config).Add(param.Config.Mi("_").Mi(cluster))
				param.Config = merger.Merge()

				var err = executor(param)
				if err != nil {
					return err
				}
			}
			return nil
		} else {
			// 1. Add data to config
			param.Config.Set("data", param.Data)
			return executor(param)
		}
	}).Input(data)
}
