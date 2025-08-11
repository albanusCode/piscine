package main

type Solver struct {
	solutionCount int
}

func (s *Solver) SolveUnique(p *Puzzle) bool {
	s.solutionCount = 0

	if !s.solveRecursive(p, 0, 0, true) {
		return false
	}

	testP := p.Clone()
	s.solutionCount = 0
	s.solveRecursive(&testP, 0, 0, false)

	return s.solutionCount == 1
}

func (s *Solver) solveRecursive(p *Puzzle, row, col int, stopAtOne bool) bool {
	if row == 9 {
		s.solutionCount++
		if stopAtOne {
			return true
		}
		return false
	}

	nextRow := row
	nextCol := col + 1
	if nextCol == 9 {
		nextCol = 0
		nextRow++
	}

	if p.grid[row][col] != 0 {
		return s.solveRecursive(p, nextRow, nextCol, stopAtOne)
	}

	for num := 1; num <= 9; num++ {
		if canPlace(p, row, col, num) {
			p.grid[row][col] = num
			if s.solveRecursive(p, nextRow, nextCol, stopAtOne) && stopAtOne {
				return true
			}
			p.grid[row][col] = 0
		}
	}

	return false
}

func canPlace(p *Puzzle, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if p.grid[row][i] == num || p.grid[i][col] == num {
			return false
		}
	}
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if p.grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}