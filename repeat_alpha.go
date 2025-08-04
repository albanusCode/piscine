package main

import "github.com/01-edu/z01"

func RepeatAlpha(s string) {
	for _, char := range s {
		for i := 'a'; i <= 'z'; i++ {
			z01.PrintRune(char)
		}
	}
}

func main() {
	RepeatAlpha("a")
}
