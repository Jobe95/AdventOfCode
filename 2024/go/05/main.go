package main

import (
	"advent-of-code-2024/utils"
	"slices"
	"strings"
)

func main() {
	utils.Run(1, partOne, 1)
	utils.Run(2, partTwo, 1)
}

type Pair struct {
	X int
	Y int
}

func partOne() int {
	output := 0
	input := utils.ReadFile("input.txt")

	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	rules := getRules(sections[0])
	updates := getUpdates(sections[1])

	for _, numbers := range updates {
		validUpdate := true
		for _, rule := range rules {
			xIndex := slices.Index(numbers, rule.X)
			yIndex := slices.Index(numbers, rule.Y)
			if xIndex != -1 && yIndex != -1 && xIndex > yIndex {
				validUpdate = false
			}
		}
		if validUpdate {
			middle := numbers[len(numbers)/2]

			output += middle
		}

	}
	return output
}

func partTwo() int {
	output := 0
	input := utils.ReadFile("input.txt")

	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	rules := getRules(sections[0])
	updates := getUpdates(sections[1])

	for _, numbers := range updates {
		var pairs []Pair
		validUpdate := true
		for _, rule := range rules {
			xIndex := slices.Index(numbers, rule.X)
			yIndex := slices.Index(numbers, rule.Y)
			if xIndex != -1 && yIndex != -1 {
				if xIndex > yIndex {
					validUpdate = false
				}
				pairs = append(pairs, rule)
			}
		}

		if !validUpdate && len(pairs) > 0 {
			for {
				correctOrder := true
				for _, rule := range pairs {
					xIndex := slices.Index(numbers, rule.X)
					yIndex := slices.Index(numbers, rule.Y)

					if xIndex > yIndex {
						correctOrder = false
						numbers = utils.MoveIndex(numbers, xIndex, yIndex)
					}
				}
				if correctOrder {
					middle := numbers[len(numbers)/2]
					output += middle
					break
				}
			}
		}
	}
	return output
}

func getUpdates(data string) [][]int {
	var updates [][]int
	lines := strings.Split(strings.TrimSpace(data), "\n")
	for _, line := range lines {
		stringNumbers := strings.Split(strings.TrimSpace(line), ",")
		updates = append(updates, utils.ToIntSlice(stringNumbers))
	}
	return updates
}

func getRules(data string) []Pair {
	var pairs []Pair
	lines := strings.Split(strings.TrimSpace(data), "\n")
	for _, line := range lines {
		parts := strings.Split(line, "|")
		pairs = append(pairs, Pair{X: utils.ParseToInt(parts[0]), Y: utils.ParseToInt(parts[1])})
	}
	return pairs
}
