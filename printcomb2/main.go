package main

import "fmt"

func main() {
	for i := '0'; i <= '9'; i++ {           // First digit of first number
		for j := '0'; j <= '9'; j++ {       // Second digit of first number
			for k := '0'; k <= '9'; k++ {   // First digit of second number
				for l := '0'; l <= '9'; l++ { // Second digit of second number
					// Compare numbers: (i,j) < (k,l)
					if i < k || (i == k && j < l) {
						fmt.Printf("%c%c %c%c", i, j, k, l)
						// Avoid trailing comma after "98 99"
						if !(i == '9' && j == '8' && k == '9' && l == '9') {
							fmt.Print(", ")
						}
					}
				}
			}
		}
	}
	fmt.Println()
}
