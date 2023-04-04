package cxcron

import (
    "time"
)

// CronJob runs the given function at the specified interval until stopped.
type CronJob struct {
    interval time.Duration
    stopChan chan struct{}
}

// NewCronJob creates a new CronJob instance with the given interval and function.
func NewCronJob(interval time.Duration, f func()) *CronJob {
    job := &CronJob{
        interval: interval,
        stopChan: make(chan struct{}),
    }

    go func() {
        f()

        ticker := time.NewTicker(job.interval)
        defer ticker.Stop()

        for {
            select {
            case <-ticker.C:
                f()
            case <-job.stopChan:
                return
            }
        }
    }()

    return job
}

// Stop stops the cron job from running.
func (job *CronJob) Stop() {
    close(job.stopChan)
}

/*

package main

import (
	"fmt"
	"time"

	"github.com/my-cron-job-package"
)

func main() {
	// Define the anonymous function to be executed by the cron job
	myFunc := func() {
		fmt.Println("Hello world!")
	}

	// Run the cron job every 5 seconds
	cronJob := cronjob.NewCronJob(myFunc, time.Second*5)

	// Start the cron job
	cronJob.Start()

	// Wait for 20 seconds
	time.Sleep(time.Second * 20)

	// Stop the cron job
	cronJob.Stop()

	// Wait for a few seconds to ensure that the cron job has stopped
	time.Sleep(time.Second * 2)
}

*/
