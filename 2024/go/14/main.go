package main

import "advent-of-code-2024/utils"

func main() {
	utils.Run(1, partOne, 1)
	utils.Run(2, partTwo, 1)
}

type Robot struct {
	DX int
	DY int
}

func partOne() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("example.txt")
	return output
}

func partTwo() int {
	output := 0
	return output
}
