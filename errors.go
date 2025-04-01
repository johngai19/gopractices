package main

import (
	"fmt"
)

func generateError() {
	// Example usage
	err := doSomething()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func doSomething() error {
	// Simulate an error
	return fmt.Errorf("an example error occurred")
}
