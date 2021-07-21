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
			if err != nil {
				p.Cache.Increase(constants.WARMUP_ERROR)
			} else {
				_ = p.Cache.UpdateFn(constants.WARMUP_DURATIONS, func(o interface{}) (interface{}, error) {
					var ms = duration.Milliseconds()
					if o == nil {
						return []int64{ms}, nil
					}

					// keep only last 1000 duration
					var queue = o.([]int64)
					if len(queue) > 1000 {
						queue[0] = 0      // assign to zero value to free memory
						queue = queue[1:] // Dequeue
					}

					queue = append(queue, ms) // Enqueue
					return queue, nil
				}, caches.Persistent)
			}
		}, duration)
	}

	return worker
}
