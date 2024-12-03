package main

import (
	"advent-of-code-2024/utils"
	"regexp"
)

func main() {
	utils.BenchmarkWithSummary("benchmark.txt", 1, partOne, 1000)
	utils.BenchmarkWithSummary("benchmark.txt", 2, partTwo, 1000)
}

func partOne() int {
	output := 0
	input := utils.ReadFile("input.txt")
	pattern := `mul\((\d+),(\d+)\)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		num1 := match[1]
		num2 := match[2]
		output += utils.ParseToInt(num1) * utils.ParseToInt(num2)
	}
	return output
}

type ActiveRanges struct {
	start int
	end   int
}

func partTwo() int {
	output := 0
	input := utils.ReadFile("input.txt")
	activePattern := `do\(\)|don't\(\)`

	reActions := regexp.MustCompile(activePattern)

	indexes := reActions.FindAllStringIndex(input, -1)
	activeRanges := []ActiveRanges{}
	previousEnabled := true
	start := 0
	for _, idx := range indexes {
		action := input[idx[0]:idx[1]]

		if !previousEnabled && action == "do()" {
			previousEnabled = true
			start = idx[1]
		} else if previousEnabled && action == "don't()" {
			activeRanges = append(activeRanges, ActiveRanges{start, idx[0]})
			previousEnabled = false
		}
	}

	if previousEnabled {
		activeRanges = append(activeRanges, ActiveRanges{start, len(input)})
	}

	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)

	for _, value := range activeRanges {
		subInput := input[value.start:value.end]
		matches := re.FindAllStringSubmatch(subInput, -1)

		for _, match := range matches {
			num1 := match[1]
			num2 := match[2]
			output += utils.ParseToInt(num1) * utils.ParseToInt(num2)
		}
	}

	return output
}
