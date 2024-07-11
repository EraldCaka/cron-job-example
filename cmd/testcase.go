package main

import "github.com/EraldCaka/cron-job-example/jobs"

func main() {
	cron := jobs.ConnectionSkeleton()
	// dummy data for the test
	cron.Customers = []*jobs.Customer{
		jobs.InitializeCustomer("customer1", 45, 125),
		jobs.InitializeCustomer("customer2", 21, 52),
		jobs.InitializeCustomer("customer3", 23, 8912),
	}
	cron.RegisterAllJobsInsideCron()
	cron.StartJobs()
	cron.Start()
	select {}
}
