package main
import "fmt"

func Count(s string) int {
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			count ++
		}
	}
	return count
}

func main() {
	x := Count("Heeaeaerdtseaawaassaew5r4eseeal2")
	fmt.Println(x)
}
