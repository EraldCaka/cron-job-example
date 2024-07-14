package main

import (
	"github.com/EraldCaka/croner/jobs"
	"github.com/EraldCaka/croner/test/customer_jobs"
)

func main() {
	cron := jobs.ConnectionSkeleton()
	customerJob := customer_jobs.CustomerConn(cron)
	// dummy data for the test
	customerJob.Customers = []*customer_jobs.Customer{
		customer_jobs.InitializeCustomer("user1", 45, 125),
		customer_jobs.InitializeCustomer("user2", 21, 52),
		customer_jobs.InitializeCustomer("user3", 23, 8912),
	}
	customerJob.RegisterAllJobsInsideCron()
	cron.StartJobs()
	cron.Start()
	defer cron.Close()
	select {}
}
