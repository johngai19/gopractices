package main

import (
	"fmt"
	"time"
)

func enumerateProducts(c1, c2 chan<- *Product) {
	for _, p := range ProductList {
		select {
		case c1 <- p:
			fmt.Println("Product sent to channel 1")
		case c2 <- p:
			fmt.Println("Product sent to channel 2")
		}
	}
	close(c1)
	close(c2)
}
func main() {
	c1 := make(chan *Product, 3)
	c2 := make(chan *Product, 2)

	go enumerateProducts(c1, c2)
	time.Sleep(1 * time.Second) // Wait for goroutine to finish

	for p := range c1 {
		fmt.Println("Received from channel 1:", p.Name)
	}

	for p := range c2 {
		fmt.Println("Received from channel 2:", p.Name)
	}
}
