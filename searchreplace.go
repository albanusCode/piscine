package main

import (
	"fmt"
	"os"
)

func searchReplace() {
	//args := os.Args takes arguments from the command-line- they come as strings
	args := os.Args
	if len(args) != 4 { // check that the strings are atleast three.
		return
	}

	str := args[1] // the first string is at index 1
	oldchar := args[2]
	newchar := args[3]

	newstr := "" // initialize an empty string newstr := "hollo"
	for _, ch := range str {
		if string(ch) == oldchar {
			newstr += newchar
		}else{
			newstr += string(ch)
		}
	}
	fmt.Println(newstr)
}
