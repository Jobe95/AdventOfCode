package utils

import (
	"fmt"
	"strconv"
)

func ParseToInt(a string) int {
	result, _ := strconv.Atoi(a)
	return result
}

func ParseToString(a int) string {
	return strconv.Itoa(a)
}

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
	if fromIndex < 0 || fromIndex > len(slice)-1 || toIndex < 0 || toIndex > len(slice)-1 {
		fmt.Println("Out of bounds!!")
		return slice
	}

	item := slice[fromIndex]
	slice = append(slice[:fromIndex], slice[fromIndex+1:]...)

	slice = append(slice[:toIndex], append([]T{item}, slice[toIndex:]...)...)
	return slice
}
