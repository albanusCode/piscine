package main
import "os"

func FirstArg() {
	if len(os.Args) < 1 {
		os.Stdout.Write([]byte("\n"))
	}
	first := os.Args[1]
	os.Stdout.Write([]byte(first + "\n"))

}

// // Write a program that takes strings as arguments, and displays its first
// argument followed by a \n.

// If the number of arguments is less than 1, the program displays \n.
