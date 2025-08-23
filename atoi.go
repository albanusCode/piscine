package main

// import (
// 	"os"
// 	"fmt"
// )

// func main(){
// 	args := os.Args
// 	if len(args) < 3 {
// 		os.Stdout.Write([]byte("Not enough arguments"))
// 	}

// 	num1 := atoi(args[1])
// 	num2 := atoi(args[2])

// 	perimeter := RectPerimeter(num1, num2)
// 	fmt.Println(perimeter)
// }

// func RectPerimeter(w, h int) int {
// 	if w < 0 || h < 0 {
// 		return -1
// 	}else{
// 		return 2*(w+h)
// 	}
// }

func atoi(str string) int {
	number := 0
	isNeg := false
	isPos := false

	for _, ch := range str {
		if ch == '-' {
			isNeg = true
		}else if ch == '+' {
			isPos = true
		}
		if ch >= '0' && ch <= '9' {
			digit := ch - '0'
			number = (number * 10) + int(digit)
		}
		if isNeg {
			return -1
		}else if isPos{
			return number * 1
		}
	}
	return number
}


// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	sum := 0

// 	args := os.Args[1]
// 	if len(os.Args) > 1 {
// 		fmt.Println(0)
// 		return
// 	}
// 	n := atoi(args)
// 	for i := 2; i <= n; i++ {
// 		isPrime := true
// 		for j := 2; j*j <= i; j++ {
// 			if i%j == 0 {
// 				isPrime = false
// 				break
// 			}
// 		}
// 		if isPrime {
// 			sum += i
// 		}
// 	}
// 	fmt.Println(sum)
// }

// func atoi(str string) int {
// 	if len(str) != 1 {
// 		return 0
// 	}

// 	isNeg := false
// 	isPos := false
// 	number := 0
// 	for _, ch := range str {
// 		if ch == '-' {
// 			isNeg = true
// 		}else if ch == '+' {
// 			isPos = true
// 		}
		
// 		if ch >= '0' && ch <= '9' {
// 			digit := ch - '0'
// 			number = (number * 10) + int(digit)
// 		}
// 		if isNeg {
// 			return number * -1
// 		}else if isPos {
// 			return number * 1
// 		}
// 	}
// 	return number
// }

