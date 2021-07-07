package schedulers

import (
	"context"
	"sync"
	"time"
)

type Scheduler struct {
	wg            *sync.WaitGroup
	cancellations []context.CancelFunc
}

func (s *Scheduler) Add(c context.Context, j Job, interval time.Duration) {
	var ctx, cancel = context.WithCancel(c)
	s.cancellations = append(s.cancellations, cancel)

	s.wg.Add(1)
	go s.process(ctx, j, interval)
}

func (s *Scheduler) Stop() {
	for _, cancel := range s.cancellations {
		cancel()
	}

	s.wg.Wait()
}

func (s *Scheduler) process(ctx context.Context, j Job, interval time.Duration) {
	var ticker = time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			j(ctx)
		case <-ctx.Done():
			s.wg.Done()
			return
		}

	}
}

func New() *Scheduler {
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}
