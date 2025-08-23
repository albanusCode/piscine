package main

import "github.com/01-edu/z01"

func AlphaBet() {
	for i := 'a'; i <= 'z'; i++ {
		if i%2 != 0 {
			z01.PrintRune(i)
		} else {
			z01.PrintRune(i-32)
		}
	}
	z01.PrintRune('\n')
}

// func main() {
// 	AlphaBet()
// }

// Write a program that displays the alphabet, with even letters in uppercase, and
// odd letters in lowercase, followed by a newline.
