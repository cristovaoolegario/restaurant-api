package main

import (
	"fmt"
	"github.com/cristovaoolegario/restaurant-api/internal/domain"
)

func main() {
	restaurant := domain.NewRestaurant(5, 3)
	go restaurant.Start()

	for i := 0; i < 20; i++ {
		o := domain.GenerateRandomOrder(i + 1)
		restaurant.OrderWg.Add(1)
		go func(order domain.Order) {
			waiter := restaurant.GetWaiterToServe(order)
			waiter.TakeOrder(order, restaurant.Orders)
		}(o)
	}

	restaurant.OrderWg.Wait()
	close(restaurant.Orders)
	fmt.Println("Restaurant closed after completing all orders.")
}
