package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Initialize a slice of integers
	nums := []int{1, 2, 3, 4, 5}

	// Print the original slice
	fmt.Println("Original slice:", nums)

	// Shuffle the slice
	shuffled := shuffle(nums)

	// Print the shuffled slice
	fmt.Println("Shuffled slice:", shuffled)
	generateError()
	stdioe()
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
