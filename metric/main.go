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
	var cmd = commandline.New(caches.Global, &models.Metadata{
		Name:    name,
		Version: version,
		Commit:  commit,
		Date:    date,
		BuiltBy: builtBy,
	}).Plugin(plugins.SupportLogLevel).
		Plugin(plugins.SupportVersion).
		Plugin(plugins.SupportListConfig).
		Plugin(plugins.SupportConfig).
		Plugin(plugins.SupportDotEnv).
		Command(&commands.Command{
			Name: commands.DEFAULT,
			Executor: func(p *commands.ExecutorParameter) error {
				conn, err := freqtrade.NewConnection(p.Config, p.Cache)
				if err != nil {
					return err
				}

				// print connection information
				p.Logger.Info(conn.String())
				// refresh cluster name
				p.Cache.Update("cluster", p.Config.Mi("freqtrade").So("cluster", "1A"), caches.Persistent)

				// start warmup
				var worker = cmd.WarmupJob(ctx, p, conn)
				// start http server
				err = cmd.HttpServer(p, conn)

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
