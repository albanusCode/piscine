package main

import (
	"fmt"
)

func ThirdTimeIsACharm(str string) string {
	length := len(str)
	if length < 1 {
		return "\n"
	}

	result := ""
	for i := 0; i < length; i++ {
		if i%3 == 2 {
			result += string(str[i])
		}
	}
	return result + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
