package main

import (
	mathutil "advent-of-code-2024"
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)


func main() {
	left, right, err := readInputs()
	if (err != nil) {
		os.Exit(3)
	}
	slices.Sort(left)
	slices.Sort(right)
	
	totalDiff := calculateDiff(left, right)
	fmt.Println("Total Difference:", totalDiff)
	
	simScore := calcSimilarityScore(left, right)
	fmt.Println("Total Similarity score:", simScore)
	
}


func readInputs() ([]int, []int, error) {
	file, err := os.Open("01/input.txt");
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close();

	var left, right []int

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
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
	count := len(left);
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
			score := mathutil.Multiply(num, count);
			similarityScore += score
		}
	}

	return similarityScore
}