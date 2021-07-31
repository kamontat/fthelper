package cmd

import (
	"context"
	"time"

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
				p.Cache.IncreaseN(constants.WARMUP_ERROR, err.Length())
			} else {
				_ = p.Cache.Bucket(
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
