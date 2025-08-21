package main

import "github.com/01-edu/z01"

func PrintZ(s string) {
	if len(s) < 1 {
		z01.PrintRune('z')
		z01.PrintRune('\n')
	}
	for _, char := range s {
		if char == 'z' {
			z01.PrintRune(char)
			z01.PrintRune('\n')
			return
		}else{
			z01.PrintRune('z')
			z01.PrintRune('\n')
		}
	}
}

// func main() {
// 	PrintZ("hello")
// }
