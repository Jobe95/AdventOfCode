package main

import (
	"advent-of-code-2024/utils"
	"sort"
)

type FileType int64

const (
	File      FileType = 0
	Free      FileType = 1
	Abandoned FileType = 2
)

type FileSystem struct {
	Type  FileType
	Value int
}

func (ft FileType) IsFreeSpace() bool {
	return ft == Free
}

func (ft FileType) IsFileSpace() bool {
	return ft == File
}

func main() {
	utils.Run(1, partOne, 20)
	utils.Run(2, partTwo, 20)

	// utils.BenchmarkWithSummary("benchmark.txt", 1, partOne, 10)
	// utils.BenchmarkWithSummary("benchmark.txt", 2, partTwo, 10)
}

func partOne() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	filesystem := []FileSystem{}

	for _, val := range input {
		id := 0
		for i, val2 := range val {
			num := int(val2 - '0')
			if i%2 == 0 {
				for i := 0; i < num; i++ {
					filesystem = append(filesystem, FileSystem{Value: id, Type: File})
				}
				id++
			} else {
				for i := 0; i < num; i++ {
					filesystem = append(filesystem, FileSystem{Value: -1, Type: Free})
				}
			}
		}
	}

	for {
		lastDigitIndex := -1
		for i := len(filesystem) - 1; i >= 0; i-- {
			if filesystem[i].Type.IsFileSpace() {
				lastDigitIndex = i
				break
			}
		}

		firstDotIndex := -1
		for i, val := range filesystem {
			if val.Type.IsFreeSpace() {
				firstDotIndex = i
				break
			}
		}
		if lastDigitIndex == -1 || firstDotIndex == -1 || firstDotIndex > lastDigitIndex {
			break
		}

		filesystem[firstDotIndex] = filesystem[lastDigitIndex]
		filesystem[lastDigitIndex] = FileSystem{Value: -1, Type: Free}
	}

	for i, val := range filesystem {
		if val.Type.IsFileSpace() {
			output += i * val.Value
		}
	}
	return output
}

func partTwo() int {
	output := 0
	input := utils.ReadFileLinesAsStringArray("input.txt")

	filesystem := []FileSystem{}

	for _, val := range input {
		id := 0
		for i, val2 := range val {
			num := int(val2 - '0')
			if i%2 == 0 {
				for i := 0; i < num; i++ {
					filesystem = append(filesystem, FileSystem{Value: id, Type: File})
				}
				id++
			} else {
				for i := 0; i < num; i++ {
					filesystem = append(filesystem, FileSystem{Value: -1, Type: Free})
				}
			}
		}
	}

	groupedCounts := make(map[int]int)

	for i := len(filesystem) - 1; i >= 0; {
		id := filesystem[i].Value

		if filesystem[i].Type.IsFreeSpace() {
			i--
			continue
		}

		count := 0

		for j := i; j >= 0 && filesystem[j].Value == id; j-- {
			count++
			i = j
		}

		groupedCounts[id] = count
		i--
	}

	keys := make([]int, 0, len(groupedCounts))
	for key := range groupedCounts {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for i := range keys {
		id := keys[i]
		if id == -2 {
			continue
		}
		count := groupedCounts[id]

		currentStart := -1
		for k := 0; k < len(filesystem); k++ {
			if filesystem[k].Value == id {
				currentStart = k
				break
			}
		}

		if currentStart == -1 {
			continue
		}

		for j := 0; j < currentStart; {
			if filesystem[j].Type.IsFreeSpace() {
				start := j
				for j < currentStart && filesystem[j].Type.IsFreeSpace() {
					j++
				}
				sequenceLength := j - start

				if sequenceLength >= count {
					for k := 0; k < len(filesystem); k++ {
						if filesystem[k].Value == id {
							filesystem[k] = FileSystem{Value: -1, Type: Abandoned}
						}
					}
					for k := 0; k < count; k++ {
						filesystem[start+k] = FileSystem{Value: id, Type: File}
					}
					break
				}
			} else {
				j++
			}
		}
	}

	for i, val := range filesystem {
		if val.Type.IsFileSpace() {
			output += i * val.Value
		}
	}
	return output
}
