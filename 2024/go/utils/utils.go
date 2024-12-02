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

func Run(part int, function func() int, runCount int) int {
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
	fmt.Printf("Part %d: \t%s%d%s \t(AVG: %s TOTAL: %s) | (Runs: %v)\n", part, green, output, reset, avg, total, runCount)
	return output
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
