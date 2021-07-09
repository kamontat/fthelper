package cmd

import (
	"context"
	"time"

	"github.com/kamontat/fthelper/metric/src/constants"
	"github.com/kamontat/fthelper/metric/src/freqtrade"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/commandline/commands"
	"github.com/kamontat/fthelper/shared/schedulers"
)

func WarmupJob(ctx context.Context, parameters *commands.ExecutorParameter, conn *freqtrade.Connection) *schedulers.Scheduler {
	var worker = schedulers.New()

	var warmupConfig = parameters.Config.Mi("warmup")

	var enabled = warmupConfig.Bo("enabled", true)
	var dur = warmupConfig.So("interval", "1m")
	var duration, _ = time.ParseDuration(dur)

	parameters.Logger.Info("warmup process: %t (interval=%s)", enabled, duration.String())
	if enabled {
		worker.Add(ctx, func(ctx context.Context) {
			parameters.Logger.Debug("warmup freqtrade connection caches")
			var duration = freqtrade.Warmup(conn).Milliseconds()

			parameters.Cache.UpdateFn(constants.WARMUP_DURATIONS, func(o interface{}) (interface{}, error) {
				if o == nil {
					return []int64{duration}, nil
				}

				// keep only last 1000 duration
				var queue = o.([]int64)
				if len(queue) > 1000 {
					queue[0] = 0      // assign to zero value to free memory
					queue = queue[1:] // Dequeue
				}

				queue = append(queue, duration) // Enqueue
				return queue, nil
			}, caches.Persistent)
		}, duration)
	}

	return worker
}
