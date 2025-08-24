package main

import (
	"fmt"
)

func Chunk(slice []int, size int) {
	if len(slice) < 1 {
		fmt.Println(slice)
		return
	}

	if size == 0 {
		fmt.Println()
		return
	}

	group := [][]int{}
	single := []int{}

	counter := 0
	for _, v := range slice {
		if counter == size {
			group = append(group, single)
			single = nil
			counter = 0
		}
		counter++
		single = append(single, v)
	}
	if counter > 0 {
		group = append(group, single)
	}
	fmt.Println(group)
}
// func main() {
// 	Chunk([]int{}, 10)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
// 	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
// }
