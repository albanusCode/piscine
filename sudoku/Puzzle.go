package main

import (
	"fmt"
)

type Puzzle struct {
	grid [9][9]int
}

func (p *Puzzle) Load(args []string) bool {
	if len(args) != 9 {
		return false
	}
	for i := 0; i < 9; i++ {
		if len(args[i]) != 9 {
			return false
		}
		for j := 0; j < 9; j++ {
			ch := args[i][j]
			if ch == '.' {
				p.grid[i][j] = 0
			} else if ch >= '1' && ch <= '9' {
				p.grid[i][j] = int(ch - '0')
			} else {
				return false
			}
		}
	}
	return p.isValid()
}

func (p *Puzzle) isValid() bool {
	for i := 0; i < 9; i++ {
		rowCheck := make(map[int]bool)
		colCheck := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if p.grid[i][j] != 0 {
				if rowCheck[p.grid[i][j]] {
					return false
				}
				rowCheck[p.grid[i][j]] = true
			}
			if p.grid[j][i] != 0 {
				if colCheck[p.grid[j][i]] {
					return false
				}
				colCheck[p.grid[j][i]] = true
			}
		}
	}

	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			boxCheck := make(map[int]bool)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					val := p.grid[boxRow*3+i][boxCol*3+j]
					if val != 0 {
						if boxCheck[val] {
							return false
						}
						boxCheck[val] = true
					}
				}
			}
		}
	}

	return true
}
func (p *Puzzle) Print() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d", p.grid[i][j])
			if j != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func (p *Puzzle) Clone() Puzzle {
	newP := Puzzle{}
	newP.grid = p.grid
	return newP
}