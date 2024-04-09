package domain

import "time"

// Dish represents a menu item
type Dish struct {
	Name            string
	Price           float64
	PreparationTime time.Duration
}
