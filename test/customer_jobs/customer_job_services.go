package customer_jobs

func (c *CustomerJob) RegisterAllJobsInsideCron() {
	c.AddJob(c.GetCustomersDataJob())
	c.AddJob(c.UpdateCustomerDataJob())
}
