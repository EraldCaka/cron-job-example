package customer_jobs

func (c *CustomerJob) RegisterAllJobsInsideCron() {
	c.AddJob(c.GetCustomersDataAndExcelExportJob())
	//c.AddJob(c.UpdateCustomerDataJob())
}
