// filepath: /workspaces/gopractices/hello_test.go
package main

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	// Input slice
	input := []int{1, 2, 3, 4, 5}

	// Call the shuffle function
	output := shuffle(input)

	// Check if the output has the same elements as the input
	if !reflect.DeepEqual(sorted(input), sorted(output)) {
		t.Errorf("Shuffle failed: expected elements %v, got %v", input, output)
	}

	// Check if the output is not in the same order as the input
	if reflect.DeepEqual(input, output) {
		t.Errorf("Shuffle failed: output is in the same order as input")
	}
}
func TestShuffleEmptySlice(t *testing.T) {
	// Test with an empty slice
	input := []int{}
	output := shuffle(input)

	// Check if the output is also an empty slice
	if len(output) != 0 {
		t.Errorf("Shuffle failed for empty slice: expected length 0, got %d", len(output))
	}
}

func TestShuffleSingleElement(t *testing.T) {
	// Test with a single-element slice
	input := []int{42}
	output := shuffle(input)

	// Check if the output is the same as the input
	if !reflect.DeepEqual(input, output) {
		t.Errorf("Shuffle failed for single-element slice: expected %v, got %v", input, output)
	}
}

func TestShuffleRandomness(t *testing.T) {
	// Test randomness by shuffling the same slice multiple times
	input := []int{1, 2, 3, 4, 5}
	shuffledOnce := shuffle(input)
	shuffledTwice := shuffle(input)

	// It's possible (but unlikely) that the two shuffles are identical
	// This test ensures that at least one shuffle produces a different order
	if reflect.DeepEqual(input, shuffledOnce) && reflect.DeepEqual(input, shuffledTwice) {
		t.Errorf("Shuffle failed to produce randomness: both shuffles are identical to the input")
	}
}

// Helper function to sort a slice
func sorted(slice []int) []int {
	sortedSlice := make([]int, len(slice))
	copy(sortedSlice, slice)
	for i := 0; i < len(sortedSlice)-1; i++ {
		for j := i + 1; j < len(sortedSlice); j++ {
			if sortedSlice[i] > sortedSlice[j] {
				sortedSlice[i], sortedSlice[j] = sortedSlice[j], sortedSlice[i]
			}
		}
	}
	return sortedSlice
}
