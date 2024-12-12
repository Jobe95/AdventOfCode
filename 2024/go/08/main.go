package main

import (
	"advent-of-code-2024/utils"
	"fmt"
)

func main() {
	utils.Run(1, partOne, 1)
	utils.Run(2, partTwo, 1)
}

func partOne() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	var grid = [][]Position{}

	for i, row := range input {
		grid = append(grid, make([]Position, len(row)))

		for j, char := range row {
			grid[i][j] = Position{X: j, Y: i, Value: int(char)}
		}
	}

	antennas := make(map[rune][]Position)

	for _, row := range grid {
		for _, pos := range row {
			char := rune(pos.Value)
			if char == '.' {
				continue
			}
			antennas[char] = append(antennas[char], pos)
		}
	}

	pairs := make(map[rune][]Pair)

	for char, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				pairs[char] = append(pairs[char], Pair{
					A: positions[i],
					B: positions[j],
				})
			}
		}
	}

	antinodes := make(map[string]bool)

	for _, pairList := range pairs {
		for _, pair := range pairList {
			extendPair(pair, len(grid[0]), len(grid), &antinodes)
		}
	}
	output = len(antinodes)
	return output
}

func partTwo() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	var grid = [][]Position{}

	for i, row := range input {
		grid = append(grid, make([]Position, len(row)))

		for j, char := range row {
			grid[i][j] = Position{X: j, Y: i, Value: int(char)}
		}
	}

	antennas := make(map[rune][]Position)

	for _, row := range grid {
		for _, pos := range row {
			char := rune(pos.Value)
			if char == '.' {
				continue
			}
			antennas[char] = append(antennas[char], pos)
		}
	}

	pairs := make(map[rune][]Pair)

	for char, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				pairs[char] = append(pairs[char], Pair{
					A: positions[i],
					B: positions[j],
				})
			}
		}
	}

	antinodes := make(map[string]bool)

	for _, positions := range antennas {
		if len(positions) > 1 {
			for _, pos := range positions {
				key := fmt.Sprintf("%d,%d", pos.X, pos.Y)
				antinodes[key] = true
			}
		}
	}

	for _, pairList := range pairs {
		for _, pair := range pairList {
			visited := map[Pair]bool{}
			extendPairRecursive(pair, len(grid[0]), len(grid), &antinodes, visited)
		}
	}
	output = len(antinodes)
	return output
}

func isInBounds(pos Position, edgeX, edgeY int) (Position, bool) {
	if (pos.X >= edgeX || pos.X < 0) || (pos.Y >= edgeY || pos.Y < 0) {
		return pos, false
	}
	return pos, true
}

func extendPairRecursive(pair Pair, edgeX, edgeY int, antinodes *map[string]bool, visited map[Pair]bool) {
	if visited[pair] {
		return
	}
	visited[pair] = true

	dx := pair.B.X - pair.A.X
	dy := pair.B.Y - pair.A.Y

	newA := Position{
		X:     pair.A.X - dx,
		Y:     pair.A.Y - dy,
		Value: '#',
	}
	newB := Position{
		X:     pair.B.X + dx,
		Y:     pair.B.Y + dy,
		Value: '#',
	}

	if posA, inBoundsA := isInBounds(newA, edgeX, edgeY); inBoundsA {
		keyA := fmt.Sprintf("%d,%d", posA.X, posA.Y)
		(*antinodes)[keyA] = true
		extendPairRecursive(Pair{newA, pair.A}, edgeX, edgeY, antinodes, visited)
	}

	if posB, inBoundsB := isInBounds(newB, edgeX, edgeY); inBoundsB {
		keyB := fmt.Sprintf("%d,%d", posB.X, posB.Y)
		(*antinodes)[keyB] = true
		extendPairRecursive(Pair{newB, pair.B}, edgeX, edgeY, antinodes, visited)
	}
}

func extendPair(pair Pair, edgeX, edgeY int, antinodes *map[string]bool) {
	dx := pair.B.X - pair.A.X
	dy := pair.B.Y - pair.A.Y

	newA := Position{
		X:     pair.A.X - dx,
		Y:     pair.A.Y - dy,
		Value: '#',
	}
	newB := Position{
		X:     pair.B.X + dx,
		Y:     pair.B.Y + dy,
		Value: '#',
	}

	if posA, inBoundsA := isInBounds(newA, edgeX, edgeY); inBoundsA {
		keyA := fmt.Sprintf("%d,%d", posA.X, posA.Y)
		(*antinodes)[keyA] = true
	}

	if posB, inBoundsB := isInBounds(newB, edgeX, edgeY); inBoundsB {
		keyB := fmt.Sprintf("%d,%d", posB.X, posB.Y)
		(*antinodes)[keyB] = true
	}
}

type Position struct {
	X     int
	Y     int
	Value int
}

type Pair struct {
	A Position
	B Position
}
