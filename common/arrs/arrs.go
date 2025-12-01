package arrs

import (
	"fmt"
	"slices"
	"strings"
)

func Contains[T comparable](arr []T, target T) bool {
	return slices.Contains(arr, target)
}

func Distinct[T comparable](arr []T) []T {
	result := make([]T, 0)
	seen := make(map[T]struct{})
	for _, val := range arr {
		if _, exists := seen[val]; !exists {
			seen[val] = struct{}{}
			result = append(result, val)
		}
	}
	return result
}

func RemoveAtIndex[T comparable](arr []T, index int) []T {
	if index < 0 || index >= len(arr) {
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}

func Join[T any](arr []T, separator string) string {
	if len(arr) == 0 {
		return ""
	}
	result := ""
	for i, val := range arr {
		if i > 0 {
			result += separator
		}
		result += fmt.Sprintf("%v", val)
	}
	return result
}

func SplitToStrings(s string, separator string) []string {
	if len(s) == 0 {
		return []string{}
	}
	return strings.Split(s, separator)
}

func SplitToInts(s string, separator string) []int {
	strs := SplitToStrings(s, separator)
	ints := make([]int, 0, len(strs))
	for _, str := range strs {
		var v int
		_, err := fmt.Sscanf(str, "%d", &v)
		if err == nil {
			ints = append(ints, v)
		}
	}
	return ints
}

func SplitToInt64s(s string, separator string) []int64 {
	strs := SplitToStrings(s, separator)
	ints := make([]int64, 0, len(strs))
	for _, str := range strs {
		var v int64
		_, err := fmt.Sscanf(str, "%d", &v)
		if err == nil {
			ints = append(ints, v)
		}
	}
	return ints
}

func SplitToFloat64s(s string, separator string) []float64 {
	strs := SplitToStrings(s, separator)
	floats := make([]float64, 0, len(strs))
	for _, str := range strs {
		var v float64
		_, err := fmt.Sscanf(str, "%f", &v)
		if err == nil {
			floats = append(floats, v)
		}
	}
	return floats
}
