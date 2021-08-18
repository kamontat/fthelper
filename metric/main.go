package main

import (
	"context"
	"os"

	"github.com/kamontat/fthelper/metric/v4/src/cmd"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/commandline"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/kamontat/fthelper/shared/commandline/models"
	"github.com/kamontat/fthelper/shared/commandline/plugins"
)

var (
	name    string = "ftmetric"
	version string = "dev"
	commit  string = "none"
	date    string = "unknown"
	builtBy string = "manually"
)

func defaultCommand(p *commands.ExecutorParameter) error {
	// start warmup
	ctx := context.Background()

	connections, err := connection.NewConnections(p.Config)
	if err != nil {
		return err
	}

	var freqtrades = freqtrade.Build(connections)
	var connectors = make([]connection.Connector, 0)
	for _, ft := range freqtrades {
		var cacheConfig = p.Config.Mi("freqtrade").Mi("cache")

		var connector = ft
		if !cacheConfig.IsEmpty() {
			// connector with cache
			connector = connection.WithCache(ft, caches.New(), cacheConfig)
		}

		// print connector information
		p.Logger.Info(connector.String())
		connectors = append(connectors, connector)
	}

	// initial connectors
	for _, connector := range connectors {
		err := connector.Initial()
		if err != nil {
			// Do panic error if initial is not success
			panic(err)
		}
	}

	// start warmup
	var worker = cmd.WarmupJob(ctx, p, connectors)
	// start http server
	err = cmd.HttpServer(p, connectors)

	// done
	worker.Stop()

	// cleanup connectors
	for _, connector := range connectors {
		err = connector.Cleanup()
		if err != nil {
			// Do panic error if cleanup is not success
			panic(err)
		}
	}
	return err
}

func main() {
	var cmd = commandline.New(caches.New(), &models.Metadata{
		Name:    name,
		Version: version,
		Commit:  commit,
		Date:    date,
		BuiltBy: builtBy,
	}).
		Plugin(plugins.SupportVersion).
		Plugin(plugins.SupportListConfig).
		Plugin(plugins.SupportDotEnv).
		Plugin(plugins.SupportConfig).
		Plugin(plugins.SupportCluster).
		Plugin(plugins.SupportLogLevel).
		Command(&commands.Command{
			Name:     commands.DEFAULT,
			Executor: defaultCommand,
		})

	var err = cmd.Start(os.Args)
	if err != nil {
		panic(err)
	}
}
