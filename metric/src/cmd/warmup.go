package cmd

import (
	"context"
	"time"

	"github.com/kamontat/fthelper/metric/v4/src/aggregators"
	"github.com/kamontat/fthelper/metric/v4/src/constants"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/kamontat/fthelper/shared/errors"
	"github.com/kamontat/fthelper/shared/schedulers"
)

func WarmupJob(ctx context.Context, p *commands.ExecutorParameter, connections []*freqtrade.Connection) *schedulers.Scheduler {
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
			for _, conn := range connections {
				p.Logger.Debug("warmup freqtrade cluster %s", conn.Cluster)
				var d, err = freqtrade.Warmup(conn)
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
	}

	return worker
}
