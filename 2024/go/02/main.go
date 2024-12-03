package main

import (
	"advent-of-code-2024/utils"
	"math"
	"strings"
)

func main() {
	utils.BenchmarkWithSummary("benchmark.txt", 1, partOne, 1000)
	utils.BenchmarkWithSummary("benchmark.txt", 2, partTwo, 1000)
}

func partOne() int {
	content := utils.ReadFileLinesAsStringArray("input.txt")
	var count int
	for _, value := range content {
		line := strings.Fields(value)
		numbers := utils.ToIntSlice(line)
		if validateLine(numbers) {
			count++
		}
	}
	return count
}

func partTwo() int {
	content := utils.ReadFileLinesAsStringArray("input.txt")
	var count int
	for _, value := range content {
		line := strings.Fields(value)
		numbers := utils.ToIntSlice(line)

		if validateLine(numbers) {
			count++
			continue
		}

		for index := range numbers {
			tmp := utils.RemoveElementImmutable(numbers, index)

			if validateLine(tmp) {
				count++
				break
			}
		}
	}
	return count
}

func validateLine(numbers []int) bool {
	var previous int
	previousDirection := 0

	for index, num := range numbers {
		if index > 0 {
			diff := num - previous
			absDiff := int(math.Abs(float64(diff)))

			if absDiff < 1 || absDiff > 3 {
				return false
			}

			currentDirection := 0
			if diff > 0 {
				currentDirection = 1
			} else if diff < 0 {
				currentDirection = -1
			}

			if previousDirection != 0 && previousDirection != currentDirection {
				return false
			}
			previousDirection = currentDirection
		}
		previous = num
	}

	return true
}
