package main

import (
	"advent-of-code-2024/utils"
)

func main() {
	utils.Run(1, partOne, 1)
	utils.Run(2, partTwo, 1)
}

func partOne() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	var grid [][]rune
	for _, line := range input {
		grid = append(grid, []rune(line))
	}

	rows := len(grid)
	cols := len(grid[0])

	patterns := []string{"XMAS", "SAMX"}

	checkPattern := func(i, j, di, dj int, pattern string) bool {
		for k := 0; k < len(pattern); k++ {
			x, y := i+k*di, j+k*dj
			if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] != rune(pattern[k]) {
				return false
			}
		}
		return true
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, pattern := range patterns {
				wordLength := len(pattern)
				// Horizontal
				if j+wordLength <= cols && checkPattern(i, j, 0, 1, pattern) {
					output++
				}
				// Vertical
				if i+wordLength <= rows && checkPattern(i, j, 1, 0, pattern) {
					output++
				}
				// Top left to bottom right
				if i+wordLength <= rows && j+wordLength <= cols && checkPattern(i, j, 1, 1, pattern) {
					output++
				}
				// Top right to bottom left
				if i+wordLength <= rows && j-wordLength >= -1 && checkPattern(i, j, 1, -1, pattern) {
					output++
				}
			}
		}
	}

	return output
}

func partTwo() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	var grid [][]rune
	for _, line := range input {
		grid = append(grid, []rune(line))
	}

	rows := len(grid)
	cols := len(grid[0])

	checkNearest := func(x1, y1, x2, y2 int) bool {
		if string(grid[x1][y1]) == "M" && string(grid[x2][y2]) == "S" || string(grid[x1][y1]) == "S" && string(grid[x2][y2]) == "M" {
			return true
		}
		return false
	}

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {

			if string(grid[i][j]) == "A" {
				tlbr := checkNearest(i-1, j-1, i+1, j+1)
				trbl := checkNearest(i-1, j+1, i+1, j-1)
				if tlbr && trbl {
					output++
				}
			}
		}
	}

	return output
}
