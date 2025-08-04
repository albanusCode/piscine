package main
import "fmt"

var d string
func Reverse(s string) {
	for i := len(s)-1; i >= 0; i -- {
		d = d + string(s[i])
	}
}

func main() {
	Reverse("Hello")
	fmt.Print(d)
}
