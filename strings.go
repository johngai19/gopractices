package main

import (
	"fmt"
	"strconv"
	"strings"
)

func processStrings() {
	grade, err := strconv.ParseFloat(strings.TrimSpace(" 3.5 "), 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return
	}
	fmt.Println("Parsed float:", grade)
}
