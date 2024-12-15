package main

import (
	"advent-of-code-2024/utils"
)

func main() {
	utils.Run(1, partOne, 1)
	utils.Run(2, partTwo, 1)
}

type Direction struct {
	Name  string
	Delta []int
}

type Position struct {
	Y int
	X int
}

func partOne() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	grid := utils.CreateGridFromLines(input)

	directions := []Direction{
		{"Right", []int{0, 1}},
		{"Down", []int{1, 0}},
		{"Left", []int{0, -1}},
		{"Up", []int{-1, 0}},
	}

	found := make(map[Position][]Position)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				visited := make(map[Position]bool)
				startPos := Position{X: j, Y: i}
				walkPath(startPos, startPos, grid, directions, visited, found)
			}
		}
	}

	for _, findings := range found {
		values := utils.UniqueValues(findings)
		output += len(values)
	}
	return output
}

func partTwo() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	grid := utils.CreateGridFromLines(input)

	directions := []Direction{
		{"Right", []int{0, 1}},
		{"Down", []int{1, 0}},
		{"Left", []int{0, -1}},
		{"Up", []int{-1, 0}},
	}

	found := make(map[Position][]Position)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				visited := make(map[Position]bool)
				startPos := Position{X: j, Y: i}
				walkPath(startPos, startPos, grid, directions, visited, found)
			}
		}
	}

	for _, findings := range found {
		output += len(findings)
	}
	return output
}

func walkPath(currentPosition Position, startPosition Position, grid [][]rune, directions []Direction, visited map[Position]bool, found map[Position][]Position) {
	visited[currentPosition] = true
	currentValue := grid[currentPosition.Y][currentPosition.X]

	for _, direction := range directions {
		y, x := currentPosition.Y+direction.Delta[0], currentPosition.X+direction.Delta[1]

		if y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
			newPosition := Position{X: x, Y: y}
			nextValue := grid[y][x]

			if visited[newPosition] || nextValue == '.' {
				continue
			}

			newValue := int(nextValue - '0')
			currentIntValue := int(currentValue - '0')

			if currentIntValue+1 == newValue {
				if newValue == 9 {
					value, exists := found[startPosition]
					if !exists {
						found[startPosition] = []Position{newPosition}
					} else {
						found[startPosition] = append(value, []Position{newPosition}...)
					}
				} else {
					walkPath(newPosition, startPosition, grid, directions, visited, found)
				}
			}
		}
	}

	visited[currentPosition] = false
}
