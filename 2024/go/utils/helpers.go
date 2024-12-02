package utils

import "strconv"

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
