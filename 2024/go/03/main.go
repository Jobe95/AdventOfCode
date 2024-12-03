package main

import (
	"advent-of-code-2024/utils"
	"regexp"
)

func main() {
	utils.Run(1, partOne, 100)
	utils.Run(2, partTwo, 100)
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

type Pattern struct {
	enabled bool
	from    int
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
	actions := []Pattern{{true, 0}}
	for _, idx := range indexes {
		action := input[idx[0]:idx[1]]
		previousEnabled := actions[len(actions)-1].enabled
		if !previousEnabled && action == "do()" {
			actions = append(actions, Pattern{true, idx[1]})
		} else if previousEnabled && action == "don't()" {
			actions = append(actions, Pattern{false, idx[1]})
		}
	}

	activeRanges := []ActiveRanges{}
	for i := 0; i < len(actions)-1; i++ {
		if actions[i].enabled {
			start := actions[i].from
			end := actions[i+1].from
			activeRanges = append(activeRanges, ActiveRanges{start, end})
		}
	}

	if actions[len(actions)-1].enabled {
		lastStart := actions[len(actions)-1].from
		lastEnd := len(input)
		activeRanges = append(activeRanges, ActiveRanges{lastStart, lastEnd})
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
