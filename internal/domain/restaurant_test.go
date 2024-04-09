package domain

import (
	"testing"
	"time"
)

func TestOrderProcessing(t *testing.T) {
	restaurant := NewRestaurant(1, 1)
	go restaurant.Start()

	restaurant.Orders <- Order{
		TableNumber: 1,
		Dishes:      []Dish{{Name: "Quick Dish", PreparationTime: 1 * time.Millisecond}},
	}
	restaurant.OrderWg.Add(1)

	doneChan := make(chan bool)
	go func() {
		restaurant.OrderWg.Wait()
		close(doneChan)
	}()

	select {
	case <-doneChan:
		break
	case <-time.After(1 * time.Second):
		t.Error("Order processing took too long")
	}
}
