package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	startTime := time.Now()
	target := rand.Intn(100) + 1
	guess := 0
	attempts := 0
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")
	fmt.Println("Enter your guess (or 0 to quit):")
	for {
		fmt.Scan(&guess)
		if guess == 0 {
			fmt.Println("Thanks for playing!")
			break
		}
		attempts++
		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number %d in %d attempts.\n", target, attempts)
			break
		}
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("You took %.2f seconds to guess the number.\n", duration.Seconds())
	fmt.Println("Thanks for playing!")
	fmt.Println("Goodbye!")
	fmt.Println("Have a great day!")
}
