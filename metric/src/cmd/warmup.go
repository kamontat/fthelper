package cmd

import (
	"context"
	"time"

	"github.com/kamontat/fthelper/metric/v4/src/aggregators"
	"github.com/kamontat/fthelper/metric/v4/src/constants"
	"github.com/kamontat/fthelper/metric/v4/src/freqtrade"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/kamontat/fthelper/shared/schedulers"
)

func WarmupJob(ctx context.Context, p *commands.ExecutorParameter, conn *freqtrade.Connection) *schedulers.Scheduler {
	var worker = schedulers.New()

	var warmupConfig = p.Config.Mi("warmup")

	var enabled = warmupConfig.Bo("enabled", true)
	var dur = warmupConfig.So("interval", "1m")
	var duration, _ = time.ParseDuration(dur)

	p.Logger.Info("warmup process: %t (interval=%s)", enabled, duration.String())
	if enabled {
		worker.Add(ctx, func(ctx context.Context) {
			p.Logger.Debug("warmup freqtrade connection caches")
			var duration, err = freqtrade.Warmup(conn)
			if err.HasError() {
				if value, ok := aggregators.Percentage(float64(err.Total()-err.Length()), float64(err.Total())); ok {
					caches.Global.Update(constants.WARMUP_SUCCEESS_RATE, value, caches.Persistent)
				}
			} else {
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
