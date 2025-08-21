package main
import "fmt"

func PrintRevAlphs() {
	for i := 'z'; i >= 'a'; i-- {
		fmt.Printf("%c ", i)
	}
	fmt.Println()
}
