package main

import (
	"advent-of-code-2024/utils"
	"fmt"
)

func main() {
	utils.Run(1, partOne, 1)
}

type MapItem int64
type Direction int64

const (
	Untouched MapItem = 0
	Obstacle  MapItem = 1
	Visited   MapItem = 2
)

const (
	Up    Direction = 0
	Right Direction = 1
	Down  Direction = 2
	Left  Direction = 3
)

type Position struct {
	X   int
	Y   int
	Dir Direction
}

func partOne() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	var grid = [][]MapItem{}
	currentPosition := Position{}

	for idx, val := range input {
		grid = append(grid, make([]MapItem, len(val)))

		for i := 0; i < len(val); i++ {
			char := val[i]
			switch char {
			case '.':
				grid[idx][i] = Untouched
			case '#':
				grid[idx][i] = Obstacle
			case '^':
				currentPosition = Position{X: idx, Y: i, Dir: Up}
			default:
				fmt.Println("Wrong!")
			}
		}
	}

	turnRight := func(dir Direction) Direction {
		return (dir + 1) % 4
	}

	isInBounds := func(x, y int, grid [][]MapItem) bool {
		return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
	}

	runMark := func(position Position) (int, int) {
		x, y := position.X, position.Y
		switch position.Dir {
		case Up:
			x--
		case Right:
			y++
		case Down:
			x++
		case Left:
			y--
		}
		return x, y
	}

	for {
		x, y := runMark(currentPosition)

		if !isInBounds(x, y, grid) {
			grid[currentPosition.X][currentPosition.Y] = Visited
			break
		}

		if grid[x][y] == Obstacle {
			currentPosition.Dir = turnRight(currentPosition.Dir)
			x, y = runMark(currentPosition)
		}
		currentPosition.X = x
		currentPosition.Y = y
		grid[x][y] = Visited
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == Visited {
				output += 1
			}
		}
	}

	return output
}
func partTwo() int {
	output := 0
	return output
}
