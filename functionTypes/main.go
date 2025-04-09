package main

import (
	"fmt"
)

type calcFunc func(float64) float64

func printPrice(product string, price float64, f calcFunc) {
	fmt.Printf("Product: %s, Price: %.2f, Price with tax: %.2f\n", product, price, f(price))
}

func priceCalcFactory(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if price > threshold {
			return price * (1 + rate)
		}
		return price
	}
}

func main() {
	watersportsProducts := map[string]float64{
		"Kayak":      275.00,
		"Canoe":      96.00,
		"Lifejacket": 48.95,
		"Hat":        15.00,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.99,
		"Shin Guards": 29.99,
		"Stadium":     75000,
	}
	waterCalc := priceCalcFactory(100.00, 0.20)
	soccerCalc := priceCalcFactory(50.00, 0.10)
	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc)
	}
}
