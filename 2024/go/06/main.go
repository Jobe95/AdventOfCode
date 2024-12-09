package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"slices"
	"strings"
)

func main() {
	// utils.BenchmarkWithSummary("benchmark.txt", 1, partOne, 1000)
	// utils.BenchmarkWithSummary("benchmark.txt", 2, partTwo, 2)
	utils.Run(1, partOne, 1)
	utils.Run(2, partTwo, 1)
}

func partOne() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	var grid = [][]MapItem{}
	currentPosition := Position{}

	for idx, val := range input {
		grid = append(grid, make([]MapItem, len(val)))

		for i := 0; i < len(grid[idx]); i++ {
			char := val[i]

			switch char {
			case '.':
				grid[idx][i] = Untouched
			case '#':
				grid[idx][i] = Obstacle
			case '^':
				currentPosition = Position{X: i, Y: idx, Dir: Up}
			default:
				fmt.Println("Wrong!")
			}
		}
	}

	for {
		x, y := movePosition(currentPosition)

		if !isInBounds(x, y, grid) {
			grid[currentPosition.Y][currentPosition.X] = Visited
			break
		}

		if grid[y][x] == Obstacle {
			currentPosition.Dir = turnRight(currentPosition.Dir)
			x, y = movePosition(currentPosition)
		}
		currentPosition.X = x
		currentPosition.Y = y
		grid[y][x] = Visited
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
	input := utils.ReadFileLinesAsStringArray("input.txt")

	var grid = [][]MapItem{}
	currentPosition := Position{}

	for idx, val := range input {
		grid = append(grid, make([]MapItem, len(val)))

		for i := 0; i < len(grid[idx]); i++ {
			char := val[i]

			switch char {
			case '.':
				grid[idx][i] = Untouched
			case '#':
				grid[idx][i] = Obstacle
			case '^':
				currentPosition = Position{X: i, Y: idx, Dir: Up}
			default:
				fmt.Println("Wrong!")
			}
		}
	}

	run := func() bool {
		position := currentPosition
		path := []Position{}
		cycles := [][]Position{}

		newGrid := make([][]MapItem, len(grid))
		for i := range grid {
			newGrid[i] = make([]MapItem, len(grid[i]))
			copy(newGrid[i], grid[i])
		}

		for {
			x, y := movePosition(position)

			if !isInBounds(x, y, newGrid) {
				return false
			}

			if newGrid[y][x] == Obstacle {
				path = append(path, position)

				counts := make(map[Position]int)
				for _, pos := range path {
					counts[pos]++
					if pos == position && counts[position] > 1 {
						startIndex := slices.Index(path, position)
						endIndex := utils.LastIndex(path, position)

						if startIndex != -1 && endIndex != -1 && startIndex != endIndex {
							cycle := append([]Position{}, path[startIndex:endIndex+1]...)

							cycles = append(cycles, cycle)
							path = append(path[:startIndex], path[endIndex+1:]...)
							path = append(path, position)

							for _, c := range cycles {
								if utils.IsSlicesEqual(c, cycle) {
									return true
								}
							}
						}
					}
				}

				position.Dir = turnRight(position.Dir)
				continue
			}
			position.X = x
			position.Y = y
			newGrid[y][x] = Visited
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == Untouched {
				grid[i][j] = Obstacle
				if run() {
					output++
				}
				grid[i][j] = Untouched
			}
		}
	}

	return output
}

func turnRight(dir Direction) Direction {
	return (dir + 1) % 4
}

func isInBounds(x, y int, grid [][]MapItem) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0])
}

func movePosition(position Position) (int, int) {
	x, y := position.X, position.Y
	switch position.Dir {
	case Up:
		y--
	case Right:
		x++
	case Down:
		y++
	case Left:
		x--
	}
	return x, y
}

func printGrid(grid [][]MapItem, currentPosition Position) {
	for y, row := range grid {
		var formattedRow []string
		for x, cell := range row {
			if x == currentPosition.X && y == currentPosition.Y {
				formattedRow = append(formattedRow, "^")
			} else {
				switch cell {
				case Untouched:
					formattedRow = append(formattedRow, "0")
				case Obstacle:
					formattedRow = append(formattedRow, "#")
				case Visited:
					formattedRow = append(formattedRow, "1")
				}
			}
		}
		fmt.Println(strings.Join(formattedRow, " "))
	}
	fmt.Println()
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
