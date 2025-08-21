package main
import "fmt"

func StringLen(s string) {
	count := 0
	for i := 0; i < len(s); i++ {
		count++
	}
	fmt.Println(count)
}

// func main() {
// 	StringLen("hello")
// }







// Write a function that returns the length of a string.
