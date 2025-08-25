package main

import (
	"fmt"
)

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}

func Itoa(x int) string {
	result := []rune{}
	isneg := false
	if x < 0 {
		isneg = true
		x = x * -1
	}
	if x == 0 {
		return "0"
	}
	runes := []rune{'0','1','2','3','4','5','6','7','8','9'}

	for x > 0 {
		digit := x % 10
		result = append([]rune{runes[digit]}, result...)
		x /= 10
	}
	if isneg {
		result = append([]rune{'-'}, result...)
	}
	return string(result)
}
