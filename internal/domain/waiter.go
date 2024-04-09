package domain

import (
	"fmt"
)

// Waiter represents a waiter in the restaurant
type Waiter struct {
	ID int
}

func (w *Waiter) TakeOrder(order Order, orderQueue chan<- Order) {
	fmt.Printf("[WAITER - %d] has taken an order for table %d\n", w.ID, order.TableNumber)
	orderQueue <- order
}

func (w *Waiter) ServeOrder(order Order) {
	fmt.Printf("[WAITER - %d] is serving the order for table %d\n", w.ID, order.TableNumber)
	fmt.Printf("[WAITER - %d] bill for table %d is %v\n", w.ID, order.TableNumber, order.CalculateBill())
}
