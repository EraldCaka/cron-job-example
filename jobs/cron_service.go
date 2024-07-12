package jobs

import (
	"fmt"
	"sync"
)

type Croner interface {
	RegisterAllJobsInsideCron()
}

func (c *Cron) Close() {
	c.Conn.Stop()
}

func (c *Cron) Start() {
	c.Conn.Start()
}

func (c *Cron) AddJob(cronExpression string, f func()) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Jobs = append(c.Jobs, &Job{CronExpression: cronExpression, Cron: f})
}

func (c *Cron) StartJobs() {
	var wg sync.WaitGroup

	c.Mu.Lock()
	defer c.Mu.Unlock()
	defer wg.Wait()

	for i, job := range c.Jobs {
		wg.Add(1)
		go func(job *Job) {
			defer wg.Done()
			_, err := c.Conn.AddFunc(job.CronExpression, job.Cron)
			if err != nil {
				fmt.Printf("Failed to start job %d: %v\n", i, err)
				return
			}
			fmt.Printf("job number %v started\n", i)
		}(job)
	}
}
