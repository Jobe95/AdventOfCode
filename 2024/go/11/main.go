package main

import (
	"advent-of-code-2024/utils"
	"strings"
)

func main() {
	utils.Run(1, partOne, 1)
	utils.Run(2, partTwo, 1)
}

func partOne() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")
	stones := utils.ToIntSlice(strings.Fields(input[0]))

	storage := make(map[int][]int)
	storage[0] = append([]int{}, stones...)

	blinks := 25

	for i := 0; i < blinks; i++ {
		for j := 0; j < len(storage[i]); j++ {
			stoneValue := storage[i][j]
			arr := checkRule(stoneValue)
			storage[i+1] = append(storage[i+1], arr...)
		}
	}

	output = len(storage[blinks])
	return output
}

func partTwo() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")
	stones := utils.ToIntSlice(strings.Fields(input[0]))

	current := make(map[int]int)
	for _, stone := range stones {
		current[stone]++
	}

	blinks := 75

	for i := 0; i < blinks; i++ {
		next := make(map[int]int)

		for stone, count := range current {
			transformed := checkRule(stone)

			for _, newStone := range transformed {
				next[newStone] += count
			}
		}
		current = next
	}

	for _, count := range current {
		output += count
	}
	return output
}

func checkRule(stone int) []int {
	if stone == 0 {
		return []int{1}
	} else if len(utils.ParseToString(stone))%2 == 0 {
		numberAsStr := utils.ParseToString(stone)
		middleIndex := len(numberAsStr) / 2
		leftStr := numberAsStr[:middleIndex]
		rightStr := numberAsStr[middleIndex:]

		left := utils.ParseToInt(leftStr)
		right := utils.ParseToInt(rightStr)

		return []int{left, right}
	}

	return []int{stone * 2024}
}
