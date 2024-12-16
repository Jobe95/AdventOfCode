package main

import (
	"advent-of-code-2024/utils"
	"sort"
)

func main() {
	utils.Run(1, partOne, 1)
	utils.Run(2, partTwo, 1)
}

type MyCell struct {
	Value   string
	Visited bool
}

func partOne() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")
	grid := utils.CreateGridFromLines(input, func(y, x int, char rune) MyCell {
		return MyCell{Value: string(char), Visited: false}
	})

	islands := make(map[string][][]utils.Cell[MyCell])

	grid.ForEach(func(cell utils.Cell[MyCell]) {
		if !cell.Item.Visited {
			island := []utils.Cell[MyCell]{}
			dfs(cell, cell, &grid, &island)
			islands[cell.Item.Value] = append(islands[cell.Item.Value], island)
		}
	})

	for _, islandList := range islands {
		for _, island := range islandList {
			perimeter := 0
			area := len(island)
			for _, cell := range island {
				perimeter += numOfNeighbours(cell, grid, utils.Directions)
			}
			output += perimeter * area
		}
	}
	return output
}

func partTwo() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")
	grid := utils.CreateGridFromLines(input, func(y, x int, char rune) MyCell {
		return MyCell{Value: string(char), Visited: false}
	})

	islands := make(map[string][][]utils.Cell[MyCell])

	grid.ForEach(func(cell utils.Cell[MyCell]) {
		if !cell.Item.Visited {
			island := []utils.Cell[MyCell]{}
			dfs(cell, cell, &grid, &island)
			islands[cell.Item.Value] = append(islands[cell.Item.Value], island)
		}
	})

	for _, islandList := range islands {
		for _, island := range islandList {
			rows := make(map[int][]utils.Cell[MyCell])
			cols := make(map[int][]utils.Cell[MyCell])
			area := len(island)
			for _, cell := range island {
				rows[cell.Y] = append(rows[cell.Y], cell)
				cols[cell.X] = append(cols[cell.X], cell)
			}

			rowUp := utils.Sum(countFences(rows, grid, utils.Up))
			colRight := utils.Sum(countFences(cols, grid, utils.Right))
			rowDown := utils.Sum(countFences(rows, grid, utils.Down))
			colLeft := utils.Sum(countFences(cols, grid, utils.Left))

			output += area * (rowUp + colRight + rowDown + colLeft)
		}
	}
	return output
}

func countFences(slice map[int][]utils.Cell[MyCell], grid utils.Grid[MyCell], direction utils.Direction) []int {
	fenceCounts := []int{}

	for _, cells := range slice {
		sort.Slice(cells, func(i, j int) bool {
			if direction == utils.Up || direction == utils.Down {
				return cells[i].X < cells[j].X
			}
			return cells[i].Y < cells[j].Y
		})
		prev := -1
		prevFence := false
		inGroup := false
		groupCount := 0
		for _, cell := range cells {
			isFence := numOfNeighbours(cell, grid, []utils.Direction{direction}) == 1
			isAdjacent := false

			switch direction {
			case utils.Up, utils.Down:
				isAdjacent = cell.X == prev+1
			case utils.Right, utils.Left:
				isAdjacent = cell.Y == prev+1
			}

			if !isAdjacent {
				inGroup = false
			}

			if !prevFence && isFence {
				inGroup = false
			}

			if !inGroup && isFence {
				groupCount++
				inGroup = true
			}
			switch direction {
			case utils.Up, utils.Down:
				prev = cell.X
			case utils.Right, utils.Left:
				prev = cell.Y
			}
			prevFence = isFence
		}
		fenceCounts = append(fenceCounts, groupCount)
	}
	return fenceCounts
}

func numOfNeighbours(cell utils.Cell[MyCell], grid utils.Grid[MyCell], directions []utils.Direction) int {
	perimeter := 0
	for _, dir := range directions {
		y, x := cell.Y+dir.DY, cell.X+dir.DX
		nextCell, inbounds := grid.Get(y, x)

		if !inbounds || nextCell.Item.Value != cell.Item.Value {
			perimeter++
		}
	}
	return perimeter
}

func dfs(currentCell utils.Cell[MyCell], previousCell utils.Cell[MyCell], grid *utils.Grid[MyCell], island *[]utils.Cell[MyCell]) {
	if currentCell.Item.Value != previousCell.Item.Value {
		return
	}

	grid.Set(currentCell.Y, currentCell.X, MyCell{Value: currentCell.Item.Value, Visited: true})
	*island = append(*island, currentCell)

	for _, dir := range utils.Directions {
		y, x := currentCell.Y+dir.DY, currentCell.X+dir.DX
		if nextCell, inBounds := grid.Get(y, x); inBounds {
			if nextCell.Item.Visited {
				continue
			}
			dfs(nextCell, currentCell, grid, island)
		}
	}
}
