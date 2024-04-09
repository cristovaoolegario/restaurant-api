package domain

import (
	"testing"
	"time"
)

func TestOrder_GetMaxPrepTime(t *testing.T) {
	tests := []struct {
		name     string
		dishes   []Dish
		expected time.Duration
	}{
		{"Single Dish", []Dish{{"Pizza", 10.0, 5 * time.Minute}},
			5 * time.Minute},
		{"Multiple Dishes", []Dish{{"Salad", 5.0, 2 * time.Minute}, {"Steak", 25.0, 10 * time.Minute}},
			10 * time.Minute},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			order := Order{TableNumber: 1, Dishes: test.dishes}
			got := order.GetMaxPrepTime()
			if got != test.expected {
				t.Errorf("Got = %v, expected %v", got, test.expected)
			}
		})
	}
}

func TestOrderCalculateBill(t *testing.T) {
	tests := []struct {
		name     string
		dishes   []Dish
		expected float64
	}{
		{"No Dishes", []Dish{}, 0},
		{"Single Dish", []Dish{{"Coffee", 3.5, 2 * time.Minute}},
			3.5},
		{"Multiple Dishes", []Dish{{"Burger", 8.0, 5 * time.Minute}, {"Fries", 2.5, 2 * time.Minute}},
			10.5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			order := Order{TableNumber: 1, Dishes: test.dishes}
			if got := order.CalculateBill(); got != test.expected {
				t.Errorf("Got = %v, expected %v", got, test.expected)
			}
		})
	}
}
