package utils

import "fmt"

func ToIntSlice(input []string) []int {
	var result []int
	for _, current := range input {
		num := ParseToInt(current)
		result = append(result, num)
	}
	return result
}

func RemoveElementImmutable(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return append([]int{}, slice...)
	}
	return append(append([]int{}, slice[:index]...), slice[index+1:]...)
}

func MoveIndex[T comparable](slice []T, fromIndex, toIndex int) []T {
	if fromIndex < 0 || fromIndex >= len(slice) || toIndex < 0 || toIndex >= len(slice) {
		fmt.Println("Out of bounds!!")
		return slice
	}

	item := slice[fromIndex]
	slice = append(slice[:fromIndex], slice[fromIndex+1:]...)

	slice = append(slice[:toIndex], append([]T{item}, slice[toIndex:]...)...)
	return slice
}

func MergeSlices[T comparable](originalSlice, newSlice []T, index int) []T {
	if index < 0 || index >= len(originalSlice) {
		panic("Index out of bounds")
	}

	before := originalSlice[:index]
	after := originalSlice[index+1:]

	merged := append(before, newSlice...)
	merged = append(merged, after...)
	return merged
}

func LastIndex[T comparable](s []T, v T) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			return i
		}
	}
	return -1
}

func IsSlicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func UniqueValues[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	unique := []T{}

	for _, val := range slice {
		if !seen[val] {
			seen[val] = true
			unique = append(unique, val)
		}
	}
	return unique
}

func Sum(slice []int) int {
	total := 0
	for _, val := range slice {
		total += val
	}
	return total
}
