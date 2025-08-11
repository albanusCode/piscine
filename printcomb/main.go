package main
import "fmt"

func main() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				if i < j && j < k {
					fmt.Printf("%c ", i)
					fmt.Printf("%c ", j)
					fmt.Printf("%c ", k)
					if !(i == '7' && j == '8' && k == '9') {
						fmt.Print(", ")
						fmt.Print(" ")
					}										
				}
			}
		}
	}
}