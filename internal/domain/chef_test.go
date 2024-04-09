package domain

import (
	"testing"
	"time"
)

func TestChef_PrepareOrder(t *testing.T) {
	chef := Chef{ID: 1}
	order := Order{
		TableNumber: 1,
		Dishes:      []Dish{{Name: "Test Dish", PreparationTime: time.Millisecond * 10}},
	}
	start := time.Now()
	chef.PrepareOrder(order)
	duration := time.Since(start)

	if duration < order.Dishes[0].PreparationTime {
		t.Errorf("Expected preparation to take at least %v, but took %v", order.Dishes[0].PreparationTime, duration)
	}
}
