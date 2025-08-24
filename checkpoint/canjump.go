package main

// import (
// 	"fmt"
// )

func CanJump(arr []uint) bool {
	if len(arr) == 0 {
		return false
	}
	if len(arr) == 1 {
		return true
	}

	position := 0
	for position < len(arr) {
		step := int(arr[position])
		if step == 0 && position != len(arr)-1 {
			return false
		}

		position += step
		if position == len(arr)-1 {
			return true
		}
		// if position >= len(arr) {
		// 	return false
		// }
	}
	return false
}

// func main() {
// 	input1 := []uint{2, 3, 1, 1, 4}
// 	fmt.Println(CanJump(input1)) // true

// 	input2 := []uint{3, 2, 1, 0, 4}
// 	fmt.Println(CanJump(input2)) // false

// 	input3 := []uint{0}
// 	fmt.Println(CanJump(input3)) // true

// 	input4 := []uint{1, 0}
// 	fmt.Println(CanJump(input4)) // true

// 	input5 := []uint{1, 1, 1, 1, 0}
// 	fmt.Println(CanJump(input5)) // true

// 	input6 := []uint{1, 1, 1, 0, 2}
// 	fmt.Println(CanJump(input6)) // false

// 	input7 := []uint{2, 2, 0, 0}
// 	fmt.Println(CanJump(input7)) // false

// 	input8 := []uint{4, 0, 0, 0, 0}
// 	fmt.Println(CanJump(input8)) // true

// 	input9 := []uint{2, 2, 2, 2, 1, 0}
// 	fmt.Println(CanJump(input9)) // true

// 	input10 := []uint{2, 1, 1, 0, 0}
// 	fmt.Println(CanJump(input10)) // false
// }
