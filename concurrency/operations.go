package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64, 3)
	for category, group := range data {
		go group.TotalPrice(category, channel)
	}
	fmt.Println("--Starting to receive results from channel--")
	for i := 0; i < len(data); i++ {
		fmt.Println("--Channel read pending")
		fmt.Println("Item buffered channel size: ", len(channel), " buffered channel capacity: ", cap(channel))
		value := <-channel
		fmt.Println("--Channel read done")
		storeTotal += value
		time.Sleep(1 * time.Second) // Simulate some processing time
	}
	fmt.Println("Total: ", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		fmt.Println("Category:", category, "product: ", p.Name, "price: ", ToCurrency(p.Price))
		total += p.Price
		time.Sleep(500 * time.Millisecond) // Simulate some processing time
	}
	fmt.Println("Category:", category, "total: ", ToCurrency(total), "channel sending")
	resultChannel <- total
	fmt.Println("Category:", category, "total: ", ToCurrency(total), "channel send complete")
}
