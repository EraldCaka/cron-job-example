package customer_jobs

import "github.com/EraldCaka/cron-job-example/jobs"

func CustomerConn(baseCron *jobs.Cron) *CustomerJob {
	return &CustomerJob{
		baseCron,
		make([]*Customer, 0),
	}
}
