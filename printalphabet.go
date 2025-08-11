package main

import "fmt"

func main() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c ", i) //%c converts the runes to strings and space after each
	}
	fmt.Println()
}
