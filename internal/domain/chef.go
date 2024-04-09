package domain

import (
	"fmt"
	"time"
)

// Chef represents a chef in the restaurant
type Chef struct {
	ID int
}

func (c *Chef) PrepareOrder(order Order) {
	fmt.Printf("[CHEF - %d] is preparing order for table %d\n", c.ID, order.TableNumber)
	time.Sleep(order.GetMaxPrepTime())
	fmt.Printf("[ORDER] for table %d is ready\n", order.TableNumber)
}
