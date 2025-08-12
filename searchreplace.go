package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return // do nothing
	}

	str := os.Args[1]
	oldChar := os.Args[2]
	newChar := os.Args[3]

	// if oldChar isn't in str, print original
	found := false
	for _, ch := range str {
		if string(ch) == oldChar {
			found = true
			break
		}
	}
	if !found {
		fmt.Println(str)
		return
	}

	// build new string manually
	result := ""
	for _, ch := range str {
		if string(ch) == oldChar {
			result += newChar
		} else {
			result += string(ch)
		}
	}

	fmt.Println(result)
}
