package main

import (
	"io"
	"os"
)

func stdioe() {
	myString := ""
	arguments := os.Args
	if len(arguments) > 1 {
		myString = arguments[1]
	} else {
		myString = "please enter a string"
		io.WriteString(os.Stderr, "string error")
		io.WriteString(os.Stderr, "\n")
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
