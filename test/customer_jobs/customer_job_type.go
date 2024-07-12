package customer_jobs

import "github.com/EraldCaka/croner/jobs"

type CustomerJob struct {
	*jobs.Cron
	Customers []*Customer
}

type Customer struct {
	Name  string
	Age   int
	Money int
}

func InitializeCustomer(name string, age int, money int) *Customer {
	return &Customer{
		Name:  name,
		Age:   age,
		Money: money,
	}
}
