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

	stones := utils.ToIntSlice(input)

	fmt.Println(input, stones)

	for _, stone := range stones {
		fmt.Println(stone)
	}

	return output
}

func partTwo() int {
	output := 0
	return output
}
