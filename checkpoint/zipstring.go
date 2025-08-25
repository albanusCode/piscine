package main

import "fmt"

func Zipstring(s string)string{
	result:=""
	count:=1
	for i:=1;i<len(s);i++{
		if s[i]==s[i-1]{
			count++
		}else{
			result+=Itoa(count)
			result+=string(s[i-1])
		}

	}
	result+=Itoa(count)
	result+=string(s[len(s)-1])
	return result
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
