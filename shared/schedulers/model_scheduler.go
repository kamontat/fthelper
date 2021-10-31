package schedulers

import (
	"context"
	"sync"
	"time"
)

// Scheduler use for schedule the job
type Scheduler struct {
	wg            *sync.WaitGroup
	cancellations []context.CancelFunc
}

// To add job and time for scheduler
func (s *Scheduler) Add(c context.Context, j Job, interval time.Duration) {
	var ctx, cancel = context.WithCancel(c)
	s.cancellations = append(s.cancellations, cancel)

	s.wg.Add(1)
	go s.process(ctx, j, interval)
}

// To stop the scheduler
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

// To create new scheduler
func New() *Scheduler {
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}
