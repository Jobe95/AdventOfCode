package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	utils.Run(1, partOne, 100)
	utils.Run(2, partTwo, 100)
}

func partOne() int {
	left, right, err := readInputs()
	if err != nil {
		panic(err)
	}

	slices.Sort(left)
	slices.Sort(right)

	return calculateDiff(left, right)
}

func partTwo() int {
	left, right, err := readInputs()
	if err != nil {
		panic(err)
	}

	slices.Sort(left)
	slices.Sort(right)

	return calcSimilarityScore(left, right)
}

func readInputs() ([]int, []int, error) {
	lines := utils.ReadFileLinesAsStringArray("input.txt")

	var left, right []int

	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])

		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("invalid numbers in line: %s", line)
		}

		left = append(left, num1)
		right = append(right, num2)
	}
	return left, right, nil
}

func calculateDiff(left []int, right []int) int {
	count := len(left)
	totalDiff := 0

	for i := 0; i < count; i++ {
		leftValue := left[i]
		rightValue := right[i]

		diff := leftValue - rightValue
		totalDiff += int(math.Abs(float64(diff)))
	}
	return totalDiff
}

func calcSimilarityScore(left, right []int) int {
	similarityScore := 0
	rightFrequency := make(map[int]int)
	for _, num := range right {
		rightFrequency[num]++
	}

	for _, num := range left {
		if count, exists := rightFrequency[num]; exists {
			score := num * count
			similarityScore += score
		}
	}

	return similarityScore
}
