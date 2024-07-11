package jobs

import (
	"fmt"
	"sync"
)

func (c *Cron) GetCustomersDataJob() (string, func()) {
	return "* * * * *", func() {
		var wg sync.WaitGroup

		c.mu.Lock()
		defer c.mu.Unlock()
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
