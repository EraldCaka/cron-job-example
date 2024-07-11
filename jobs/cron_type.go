package jobs

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type Customer struct {
	Name  string
	Age   int
	Money int
}

type Job struct {
	CronExpression string
	Cron           func()
}

type Cron struct {
	Conn      *cron.Cron
	mu        sync.Mutex
	Customers []*Customer
	Jobs      []*Job
}

func InitializeCustomer(name string, age int, money int) *Customer {
	return &Customer{
		Name:  name,
		Age:   age,
		Money: money,
	}
}
