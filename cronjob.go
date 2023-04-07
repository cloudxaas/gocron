package cxcron

import (
	"sync"
	"time"
)

// CronJob runs the given function at the specified interval until stopped.
type CronJob struct {
    interval time.Duration
    stopChan chan struct{}
    f        func()
    ticker   *time.Ticker
    wg       sync.WaitGroup
}

// NewCronJob creates a new CronJob instance with the given interval and function.
func NewCronJob(interval time.Duration, f func()) *CronJob {
    job := &CronJob{
        interval: interval,
        stopChan: make(chan struct{}),
        f:        f,
    }

    job.wg.Add(1)
    go job.run()

    return job
}

// run starts the cron job.
func (job *CronJob) run() {
    defer job.wg.Done()

    // Call the function immediately before starting the ticker
    job.f()

    job.ticker = time.NewTicker(job.interval)
    defer job.ticker.Stop()

    for {
        select {
        case <-job.ticker.C:
            job.f()
        case <-job.stopChan:
            return
        }
    }
}

// Stop stops the cron job from running.
func (job *CronJob) Stop() {
    close(job.stopChan)
    job.wg.Wait()
}