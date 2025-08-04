package main

import "github.com/01-edu/z01"

func DigitOrder() {

	for i := '9'; i >= '0'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

func main() {
	DigitOrder()
}

// Write a program that displays all digits in descending order, followed by a
// newline.
