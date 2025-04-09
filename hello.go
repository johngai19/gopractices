package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	// Initialize a slice of integers
	nums := []int{1, 2, 3, 4, 5}

	// Print the original slice
	fmt.Println(strings.ToLower("Original slice:"), nums)

	// Shuffle the slice
	shuffled := shuffle(nums)

	// Print the shuffled slice
	fmt.Println(strings.ToTitle("Shuffled slice:"), shuffled)
	// Print the original slice again to show it remains unchanged
	generateError()
	//stdioe()
	stdioe()
	//datetimes()
	datetimes()
	//processStrings()
	processStrings()
	//funcs()
	products := map[string]float64{
		"item1": 50.0,
		"item2": 150.0,
		"item3": 200.0,
	}
	total, tax := calcTotalPrice(products, 0)
	fmt.Println("Total price:", total)
	fmt.Printf("Total tax: %.2f\n", tax)

}

func shuffle(slice []int) []int {
	// Create a copy of the slice to avoid modifying the original
	shuffled := make([]int, len(slice))
	copy(shuffled, slice)

	// Shuffle the copied slice using Fisher-Yates algorithm
	for i := len(shuffled) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled
}
