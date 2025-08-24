package main

// import "fmt"

func CheckNumber(str string) bool {
	for _, ch := range str {
		if ch >= '0' && ch <= '9' {
			return true
		}
	}
	return false
}

// func main() {
// 	fmt.Println(CheckNumber("Hello"))
// 	fmt.Println(CheckNumber("Hello1"))
// }
