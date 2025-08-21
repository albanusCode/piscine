package main

import "github.com/01-edu/z01"

func PrintScr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

// func main() {
// 	PrintScr("Hello World!")
// }

// Write a program that displays "Hello World!" followed by a \n.
