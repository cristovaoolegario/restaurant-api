package domain

import "sync"

// Restaurant encapsulates the restaurant operations
type Restaurant struct {
	Chefs   []*Chef
	Waiters []*Waiter
	Orders  chan Order
	OrderWg sync.WaitGroup
}

func NewRestaurant(numChefs, numWaiters int) *Restaurant {
	chefs := make([]*Chef, numChefs)
	for i := range chefs {
		chefs[i] = &Chef{ID: i + 1}
	}
	waiters := make([]*Waiter, numWaiters)
	for i := range waiters {
		waiters[i] = &Waiter{ID: i + 1}
	}
	return &Restaurant{
		Chefs:   chefs,
		Waiters: waiters,
		Orders:  make(chan Order, 100),
	}
}

func (r *Restaurant) Start() {
	for _, chef := range r.Chefs {
		go func(c *Chef) {
			for order := range r.Orders {
				c.PrepareOrder(order)
				r.OrderWg.Done()
				for _, waiter := range r.Waiters {
					if waiter.ID == r.GetWaiterIdToTable(order) {
						waiter.ServeOrder(order)
						break
					}
				}
			}
		}(chef)
	}
}

func (r *Restaurant) GetWaiterToServe(order Order) *Waiter {
	return r.Waiters[r.GetWaiterIdToTable(order)]
}

func (r *Restaurant) GetWaiterIdToTable(order Order) int {
	// It's a way of each waiter get a table
	return order.TableNumber % len(r.Waiters)
}
