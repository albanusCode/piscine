package main

import (
	"fmt"
	"os"
)

func main() {
	p := Puzzle{}
	if !p.Load(os.Args[1:]) {
		fmt.Println("Error")
		return
	}

	s := Solver{}
	if !s.SolveUnique(&p) {
		fmt.Println("Error")
		return
	}

	p.Print()
}