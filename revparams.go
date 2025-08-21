package main

import (
	"fmt"
	"os"
)

func revParams() {
	args := os.Args[1:] // skip the program name

	// loop in reverse order
	for i := len(args) - 1; i >= 0; i-- {
		fmt.Println(args[i])
	}
}
