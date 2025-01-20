package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"apple", "banana", "cherry", "date", "fig"}
	target := "cherry"

	index := sort.Search(len(strs), func(i int) bool {
		return strs[i] >= target
	})

	if index < len(strs) && strs[index] == target {
		fmt.Printf("Элемент '%s' найден на индексе %d\n", target, index)
	} else {
		fmt.Printf("Элемент '%s' не найден\n", target)
	}

	// index := binarySearch(arr, target)
	// if index != -1 {
	// 	fmt.Printf("Элемент %d найден на индексе %d\n", target, index)
	// } else {
	// 	fmt.Printf("Элемент %d не найден\n", target)
	// }
}

// вручную
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
