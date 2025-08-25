package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	var freq1 [256]int
	var freq2 [256]int

	// Mark presence of each character in str1
	for i := 0; i < len(str1); i++ {
		freq1[str1[i]] = 1
	}

	// Mark presence of each character in str2
	for i := 0; i < len(str2); i++ {
		freq2[str2[i]] = 1
	}

	count := 0

	// Check for characters present in only one of the strings
	for i := 0; i < 256; i++ {
		if freq1[i] != freq2[i] {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
