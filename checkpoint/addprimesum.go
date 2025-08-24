package main

import (
	"fmt"
	"os"
)

func main() {
	sum := 0

	args := os.Args[1]

	if len(os.Args[1:]) > 1 {
		fmt.Println(0)
		return
	}

	n := atoi(args)
	for i := 2; i <= n; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			sum += i
		}
	}
	fmt.Println(sum)
}

func atoi(str string) int {
	number := 0

	isNeg := false
	isPos := false

	for _, ch := range str {
		if ch == '-' {
			isNeg = true
		} else if ch == '+' {
			isPos = true
		}
		if ch >= '0' && ch <= '9' {
			digit := ch - '0'
			number = (number * 10) + int(digit)
		}
		if isNeg {
			return number * -1
		} else if isPos {
			return number * 1
		}
	}
	return number
}
