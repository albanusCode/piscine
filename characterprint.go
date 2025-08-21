package main
// import "os"

import "github.com/01-edu/z01"

func CharPrint(s string) {

	for _, char := range s {
		if char == 'a' {
			z01.PrintRune('a')
			z01.PrintRune('\n')
			return
		}
	}
	z01.PrintRune('\n')
}

// func main() {
// 	slice := os.Args
// 	if len(slice) == 2 {
// 		CharPrint(slice[1])
// 	} else {
// 		z01.PrintRune('a')
// 		z01.PrintRune('\n')
// 	}
// }


// Write a program that takes a string, and displays the first 'a' character it
// encounters in it, followed by a newline. If there are no 'a' characters in the
// string, the program just writes a newline. If the number of parameters is not
// 1, the program displays 'a' followed by a newline.
