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

	grid := utils.CreateGridFromLines(input, func(y, x int, char rune) rune {
		return char
	})

	found := make(map[utils.Cell[rune]][]utils.Cell[rune])

	grid.ForEach(func(cell utils.Cell[rune]) {
		if cell.Item == '0' {
			visited := make(map[utils.Cell[rune]]bool)
			walkPath(cell, cell, grid, utils.Directions, visited, found)
		}
	})

	for _, findings := range found {
		values := utils.UniqueValues(findings)
		output += len(values)
	}
	return output
}

func partTwo() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	grid := utils.CreateGridFromLines(input, func(y, x int, char rune) rune {
		return char
	})

	found := make(map[utils.Cell[rune]][]utils.Cell[rune])

	grid.ForEach(func(cell utils.Cell[rune]) {
		if cell.Item == '0' {
			visited := make(map[utils.Cell[rune]]bool)
			walkPath(cell, cell, grid, utils.Directions, visited, found)
		}
	})

	for _, findings := range found {
		output += len(findings)
	}
	return output
}

func walkPath(currentCell utils.Cell[rune], startCell utils.Cell[rune], grid utils.Grid[rune], directions []utils.Direction, visited map[utils.Cell[rune]]bool, found map[utils.Cell[rune]][]utils.Cell[rune]) {
	visited[currentCell] = true

	for _, direction := range directions {
		y, x := currentCell.Y+direction.DY, currentCell.X+direction.DX

		if nextCell, inBounds := grid.Get(y, x); inBounds {
			if visited[nextCell] || nextCell.Item == '.' {
				continue
			}
			newValue := int(nextCell.Item - '0')
			currentIntValue := int(currentCell.Item - '0')

			if currentIntValue+1 == newValue {
				if newValue == 9 {
					value, exists := found[startCell]
					if !exists {
						found[startCell] = append(found[startCell], nextCell)
					} else {
						found[startCell] = append(value, nextCell)
					}
				} else {
					walkPath(nextCell, startCell, grid, directions, visited, found)
				}
			}
		}
	}

	visited[currentCell] = false
}
