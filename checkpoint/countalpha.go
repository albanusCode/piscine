package main

import "fmt"

func CountAlpha(str string) int {
	count := 0
	for _, ch := range str {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			count++
		}
	}
	return count
}

// func main() {
// 	fmt.Println(CountAlpha("Hello world"))
// 	fmt.Println(CountAlpha("H e l l o"))
// 	fmt.Println(CountAlpha("H1e2l3l4o"))
// }
