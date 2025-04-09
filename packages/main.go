package main

import (
	"fmt"
	"packages/store"
	"packages/store/cart"

	"github.com/fatih/color"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279.00)
	fmt.Printf("Product: %s\n", product.Name)
	fmt.Printf("Price: $%.2f\n", product.Price())
	cart := cart.Cart{
		CustomerName: "John Doe",
		Products:     []store.Product{*product},
	}
	color.Green("Cart Details:")
	color.Blue("Customer: %s", cart.CustomerName)
	fmt.Printf("Customer: %s\n", cart.CustomerName)
	fmt.Printf("Total: $%.2f\n", cart.GetTotal())
}
