package customer_jobs

import (
	"fmt"
	"sync"
)

func (c *CustomerJob) GetCustomersDataJob() (string, func()) {
	return "* * * * *", func() {
		var wg sync.WaitGroup
		fmt.Println("get customers data job")
		c.Mu.Lock()
		defer c.Mu.Unlock()
		defer wg.Wait()
		for _, customer := range c.Customers {
			wg.Add(1)
			go func(customer *Customer) {
				defer wg.Done()
				fmt.Printf("%s has %d money\n", customer.Name, customer.Money)
			}(customer)
		}
	}
}
func (c *CustomerJob) UpdateCustomerDataJob() (string, func()) {
	return "* * * * *", func() {
		var wg sync.WaitGroup
		fmt.Println("update customers data job")
		c.Mu.Lock()
		defer c.Mu.Unlock()
		defer wg.Wait()
		for _, customer := range c.Customers {
			wg.Add(1)
			go func(customer *Customer) {
				defer wg.Done()
				customer.Money += 1000

				fmt.Printf("%s has %d money\n", customer.Name, customer.Money)
			}(customer)
		}
	}
}
