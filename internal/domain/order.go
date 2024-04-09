package domain

import (
	"math/rand"
	"time"
)

// Order represents a customer's order
type Order struct {
	TableNumber int
	Dishes      []Dish
}

func (o *Order) GetMaxPrepTime() time.Duration {
	maxPrepTime := time.Duration(0)
	for _, dish := range o.Dishes {
		if dish.PreparationTime > maxPrepTime {
			maxPrepTime = dish.PreparationTime
		}
	}
	return maxPrepTime
}

func (o *Order) CalculateBill() float64 {
	var total float64
	for _, item := range o.Dishes {
		total += item.Price
	}
	return total
}

func GenerateRandomOrder(tableNumber int) Order {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	dishes := []Dish{
		{Name: "Pizza", PreparationTime: 5 * time.Second, Price: 15.5},
		{Name: "Pasta", PreparationTime: 3 * time.Second, Price: 25.3},
		{Name: "Burger", PreparationTime: 4 * time.Second, Price: 13.5},
		{Name: "IceCream", PreparationTime: 2 * time.Second, Price: 2.5},
	}

	numDishes := rand.Intn(len(dishes))
	var orderDishes []Dish
	for i := 0; i <= numDishes; i++ {
		dishIndex := rand.Intn(len(dishes))
		orderDishes = append(orderDishes, dishes[dishIndex])
	}

	return Order{TableNumber: tableNumber, Dishes: orderDishes}
}
