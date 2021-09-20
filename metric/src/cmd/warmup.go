package cmd

import (
	"context"
	"time"

	"github.com/kamontat/fthelper/metric/v4/src/aggregators"
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/constants"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/kamontat/fthelper/shared/errors"
	"github.com/kamontat/fthelper/shared/schedulers"
)

func WarmupJob(ctx context.Context, p *commands.ExecutorParameter, connectors []connection.Connector) *schedulers.Scheduler {
	var worker = schedulers.New()

	var warmupConfig = p.Config.Mi("warmup")

	var enabled = warmupConfig.Bo("enabled", true)
	var dur = warmupConfig.So("interval", "1m")
	var duration, _ = time.ParseDuration(dur)

	p.Logger.Info("warmup process: %t (interval=%s)", enabled, duration.String())
	if enabled {
		worker.Add(ctx, func(ctx context.Context) {
			var globalError = errors.New()
			var duration = time.Duration(0)

			for _, connector := range connectors {
				p.Logger.Debug("warmup freqtrade cluster %s", connector.Cluster())

				var start = time.Now()

				// Warmup connection with freqtrade (APIs and DB)
				err := connector.ConnectAll()

				var d = time.Since(start)

				if value, ok := aggregators.Percentage(float64(err.Total()-err.Length()), float64(err.Total())); ok {
					caches.Global.Update(constants.WARMUP_SUCCEESS_RATE, value, caches.Persistent)
				}

				duration = duration + d
				globalError.Merge(err)
			}

			if !globalError.HasError() {
				_ = caches.Global.Bucket(
					constants.WARMUP_DURATIONS,
					int64(duration.Milliseconds()),
					1000,
					caches.Persistent,
				)
			}
		}, duration)
	} else {
		// If warmup is disabled, the success rate will be negative number
		caches.Global.Update(constants.WARMUP_SUCCEESS_RATE, -1, caches.Persistent)
	}

	return worker
}
