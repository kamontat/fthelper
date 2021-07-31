package main

import (
	"context"
	"os"

	"github.com/kamontat/fthelper/metric/v4/src/cmd"
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

func main() {
	// start warmup
	ctx := context.Background()
	var cmd = commandline.New(caches.New(), &models.Metadata{
		Name:    name,
		Version: version,
		Commit:  commit,
		Date:    date,
		BuiltBy: builtBy,
	}).Plugin(plugins.SupportLogLevel).
		Plugin(plugins.SupportVersion).
		Plugin(plugins.SupportListConfig).
		Plugin(plugins.SupportDotEnv).
		Plugin(plugins.SupportConfig).
		Command(&commands.Command{
			Name: commands.DEFAULT,
			Executor: func(p *commands.ExecutorParameter) error {
				connections, err := freqtrade.NewConnections(p.Config)
				if err != nil {
					return err
				}

				for _, conn := range connections {
					// print connection information
					p.Logger.Info(conn.String())
				}

				// start warmup
				var worker = cmd.WarmupJob(ctx, p, connections)
				// start http server
				err = cmd.HttpServer(p, connections)

				// done
				worker.Stop()
				return err
			},
		})

	var err = cmd.Start(os.Args)
	if err != nil {
		panic(err)
	}
}
