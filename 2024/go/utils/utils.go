package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(data))
}

func ReadFileLinesAsStringArray(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func Run(part int, function func() int, runCount int) (int, time.Duration, time.Duration, int) {
	var durations []time.Duration
	var output int

	for i := 0; i < runCount; i++ {
		start := time.Now()
		output = function()
		duration := time.Since(start)
		durations = append(durations, duration)
	}

	avg, total := calculateAverageAndTotalDuration(durations)

	green := "\033[32m"
	reset := "\033[0m"
	fmt.Printf("Part %d: \t%s%d%s \tAVG: %s\n", part, green, output, reset, avg)
	return output, avg, total, len(durations)
}

func BenchmarkWithSummary(fileName string, part int, function func() int, runCount int) {
	_, avg, total, runs := Run(part, function, runCount)

	err := writeBenchmarkSummary(fileName, part, avg, total, runs)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	}
}

func calculateAverageAndTotalDuration(durations []time.Duration) (time.Duration, time.Duration) {
	if len(durations) == 0 {
		return 0, 0
	}
	var total time.Duration
	for _, d := range durations {
		total += d
	}
	return total / time.Duration(len(durations)), total
}

func writeBenchmarkSummary(fileName string, part int, avg, total time.Duration, runs int) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if part == 2 {
		_, err = file.WriteString(fmt.Sprintf(
			"Part %d - Average: %s, Total: %s, Runs: %d\n\n",
			part, avg, total, runs,
		))
	} else {
		_, err = file.WriteString(fmt.Sprintf(
			"Part %d - Average: %s, Total: %s, Runs: %d\n",
			part, avg, total, runs,
		))
	}
	if err != nil {
		return err
	}

	return nil
}
