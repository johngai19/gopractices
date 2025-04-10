package main

import (
	"fmt"
	"math/rand"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{
	"John Doe",
	"Jane Smith",
	"Bob Johnson",
	"Emily Davis",
	"Michael Brown",
	"Sarah Wilson",
	"David Garcia",
}

func DispatchOrders(channel chan<- DispatchNotification) {
	orderCount := rand.Intn(20) + 2
	fmt.Println("Order count: ", orderCount)
	fmt.Println("Dispatching orders to channel")
	fmt.Println("Channel size: ", len(channel), " Channel capacity: ", cap(channel))
	fmt.Println("Channel buffered size: ", cap(channel))
	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Quantity: rand.Intn(10),
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
		}

	}
	close(channel)
}
