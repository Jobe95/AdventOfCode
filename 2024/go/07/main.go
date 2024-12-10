package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strings"
)

func main() {
	utils.Run(1, partOne, 1)
	utils.Run(2, partTwo, 1)
}

const (
	Multiplication int = 0
	Addition       int = 1
	Concetenation  int = 2
)

func evalCombination(values []int, target int, currentIndex int, currentResult int, operators []int) (int, bool) {
	if currentIndex == len(values) {
		if currentResult == target {
			return currentResult, true
		}
		return currentResult, false
	}

	currentNum := values[currentIndex]

	for _, operator := range operators {
		var result int

		switch operator {
		case Addition:
			result = currentResult + currentNum
		case Multiplication:
			result = currentResult * currentNum
		case Concetenation:
			concatResult := utils.ParseToInt(fmt.Sprintf("%d%d", currentResult, currentNum))
			result = concatResult
		}
		if result, match := evalCombination(values, target, currentIndex+1, result, operators); match {
			return result, true
		}

	}
	return currentResult, false
}

func partOne() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	operators := []int{Multiplication, Addition}

	for _, line := range input {
		parts := strings.Split(line, ":")
		total := utils.ParseToInt(parts[0])
		values := utils.ToIntSlice(strings.Fields(parts[1]))

		result, match := evalCombination(values, total, 1, values[0], operators)
		if match {
			output += result
		}
	}

	return output
}
func partTwo() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	operators := []int{Multiplication, Addition, Concetenation}

	for _, line := range input {
		parts := strings.Split(line, ":")
		total := utils.ParseToInt(parts[0])
		values := utils.ToIntSlice(strings.Fields(parts[1]))

		result, match := evalCombination(values, total, 1, values[0], operators)
		if match {
			output += result
		}
	}

	return output
}
