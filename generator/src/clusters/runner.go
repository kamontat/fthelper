package clusters

import (
	"github.com/kamontat/fthelper/shared/configs"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/runners"
)

func NewRunnerV2(data, mapper maps.Mapper, executor Executor, setting *Settings) *runners.Runner {
	var name = data.So("display", "default")
	var config = mapper.Copy()
	var index = GetIndex()
	var generatorType = data.Si("type")
	var logger = loggers.Get("generator", name)
	// add .clusters to data
	var raws = config.Ai("clusters")
	if len(raws) < 1 {
		raws = append(raws, "") // add default empty cluster
	}

	var parameter = &ExecutorParameter{
		Index:     index,
		Name:      name,
		Type:      generatorType,
		Config:    config,
		VarConfig: config.Mi("variables"),
		Logger:    logger,
		Data:      data.Set("clusters", raws),
	}

	return runners.NewRunner(
		name,
		runners.NoValidator,
		func(context *runners.Context) error {
			var param = context.Input().(*ExecutorParameter)
			param.Logger.Info("Start generate %s (%s) with data: %v", param.Name, param.Type, param.Data)
			if param.Data.Bo("withCluster", setting.DefaultWithCluster) {
				for _, raw := range param.Data.Ai("clusters") {
					var cluster = raw.(string)
					param.Logger.Debug("override data with cluster: %s", cluster) // this might be remove because cluster never changes

					// 1. Add cluster to data
					param.Data.Set("cluster", cluster)
					// 2. Add data to config
					param.Config.Set("data", param.Data)
					// 3. Add cluster name to variables, for in file system template
					param.Config.Set("variables.cluster", cluster)
					// 4. Override config with cluster
					param.Config = configs.BuildClusterConfig(cluster, param.Config)

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
		},
	).
		Disable(parameter.Data.Bo("disabled", false)).
		Input(parameter)
}
