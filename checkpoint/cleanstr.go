package main

import (
	"fmt"
)

func CleanStr(str string) {
	if len(str) <=1 {
		fmt.Println()
		return
	}

	output := ""
	for i := 0; i < len(str); i++ {
		if i > 0 && str[i] == ' ' && str[i-1] == ' ' {
			continue
		}
		output += string(str[i])
	}

	result := ""
	for i := 0; i < len(output); i++ {
		if (i == 0 && output[i] == ' ') || (i == len(output)-1 && output[i] == ' ') {
			continue
		}
		result += string(output[i])
	}

	fmt.Println(result)
}

func main() {
	CleanStr("you see it's easy to display the same thing")
	CleanStr(" only    it's  harder   ")
	CleanStr(" how funny")
	CleanStr("Did you   hear Mathilde ?")
	CleanStr("")
}
