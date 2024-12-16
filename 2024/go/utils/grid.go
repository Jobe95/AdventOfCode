package utils

import "fmt"

type Cell[T comparable] struct {
	Y, X int
	Item T
}

type Grid[T comparable] struct {
	Cells [][]Cell[T]
}

type Direction struct {
	Name   string
	DY, DX int
}

var (
	Right = Direction{Name: "Right", DY: 0, DX: 1}
	Down  = Direction{Name: "Down", DY: 1, DX: 0}
	Left  = Direction{Name: "Left", DY: 0, DX: -1}
	Up    = Direction{Name: "Up", DY: -1, DX: 0}
)

var Directions = []Direction{Right, Down, Left, Up}

func CreateGridFromLines[T comparable](input []string, initializer func(y, x int, char rune) T) Grid[T] {
	cells := make([][]Cell[T], len(input))

	for y, row := range input {
		cells[y] = make([]Cell[T], len(row))
		for x, char := range row {
			cells[y][x] = Cell[T]{
				Y:    y,
				X:    x,
				Item: initializer(y, x, char),
			}
		}
	}
	return Grid[T]{Cells: cells}
}

func CreateGridFromDimensions[T comparable](rows, cols int, initializer func(y, x int) T) Grid[T] {
	cells := make([][]Cell[T], rows)

	for y := 0; y < rows; y++ {
		cells[y] = make([]Cell[T], cols)
		for x := 0; x < cols; x++ {
			cells[y][x] = Cell[T]{
				Y:    y,
				X:    x,
				Item: initializer(y, x),
			}
		}
	}
	return Grid[T]{Cells: cells}
}

func (g Grid[T]) Get(y, x int) (Cell[T], bool) {
	if IsInBounds(g, y, x) {
		return g.Cells[y][x], true
	}

	return Cell[T]{}, false
}

func (g *Grid[T]) Set(y, x int, value T) {
	if IsInBounds(*g, y, x) {
		g.Cells[y][x].Item = value
	} else {
		fmt.Printf("Attempted to Set out-of-bounds cell (Y=%d, X=%d)\n", y, x)
	}
}

func (g Grid[T]) Dimensions() (rows, cols int) {
	if len(g.Cells) == 0 {
		return 0, 0
	}
	return len(g.Cells), len(g.Cells[0])
}

func (g Grid[T]) ForEach(action func(cell Cell[T])) {
	for _, row := range g.Cells {
		for _, cell := range row {
			action(cell)
		}
	}
}

func (g Grid[T]) PrintGrid() {
	for _, row := range g.Cells {
		for _, cell := range row {
			fmt.Printf("[%v (Y=%d, X=%d)] ", cell.Item, cell.Y, cell.X)
		}
		fmt.Println()
	}
}

func IsInBounds[T comparable](grid Grid[T], y, x int) bool {
	return y >= 0 && y < len(grid.Cells) && x >= 0 && x < len(grid.Cells[y])
}
