package main

import (
	"fmt"
)

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	expenses := []Expense{
		Product{"Kayak", "Watercraft", 275.00},
		Service{"Boat Cover", 12, 89.5},
		&Service{"Boat Insurance", 12, 100.0},
		&Product{"Lifejacket", "Safety Gear", 48.95},
	}
	for _, expense := range expenses {
		switch e := expense.(type) {
		case *Service:
			fmt.Printf("*Service: %s, Monthly Cost: %.2f\n", e.getName(), e.getCost(false))
		case *Product:
			fmt.Printf("*Product: %s, Annual Cost: %.2f\n", e.getName(), e.getCost(true))
		case Service:
			fmt.Printf("Service: %s, Monthly Cost: %.2f\n", e.getName(), e.getCost(false))
		case Product:
			fmt.Printf("Product: %s, Annual Cost: %.2f\n", e.getName(), e.getCost(true))
		default:
			fmt.Printf("Unknown type: %T\n", e)
		}
	}
}
