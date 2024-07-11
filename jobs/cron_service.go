package jobs

import (
	"fmt"
	"sync"
)

func (c *Cron) Close() {
	c.Conn.Stop()
	c.Conn.Entries()
}

func (c *Cron) Start() {
	c.Conn.Start()
}

func (c *Cron) AddJob(cronExpression string, f func()) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Jobs = append(c.Jobs, &Job{CronExpression: cronExpression, Cron: f})
}

func (c *Cron) RegisterAllJobsInsideCron() {
	// tried many instances of the same job so i could test the concurrency of all the jobs
	c.AddJob(c.GetCustomersDataJob())
	c.AddJob(c.GetCustomersDataJob())
	c.AddJob(c.GetCustomersDataJob())
	c.AddJob(c.GetCustomersDataJob())
	c.AddJob(c.GetCustomersDataJob())
}

func (c *Cron) StartJobs() {
	var wg sync.WaitGroup

	c.mu.Lock()
	defer c.mu.Unlock()
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
