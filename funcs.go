package main

import (
	"fmt"
)

func calcTax(price float64, taxRate float64) (float64, bool) {

	if price > 100 {
		return price * taxRate, true
	}
	return 0, false
}

func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
	total = minSpend
	tax = 0
	for product, price := range products {
		if taxAmount, due := calcTax(price, 0.1); due {
			total += price + taxAmount
			tax += taxAmount
		} else {
			total += price
		}
		fmt.Printf("Product %s: price %.2f, tax %.2f, Total price is %.2f\n", product, price, tax, total)

	}
	return
}
