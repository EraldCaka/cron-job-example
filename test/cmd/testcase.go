package main

import (
	"github.com/EraldCaka/cron-job-example/jobs"
	"github.com/EraldCaka/cron-job-example/test/customer_jobs"
)

func main() {
	cron := jobs.ConnectionSkeleton()
	customerJob := customer_jobs.CustomerConn(cron)
	// dummy data for the test
	customerJob.Customers = []*customer_jobs.Customer{
		customer_jobs.InitializeCustomer("customer1", 45, 125),
		customer_jobs.InitializeCustomer("customer2", 21, 52),
		customer_jobs.InitializeCustomer("customer3", 23, 8912),
	}
	customerJob.RegisterAllJobsInsideCron()
	cron.StartJobs()
	cron.Start()
	defer cron.Close()
	select {}
}
