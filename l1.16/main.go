package main

import (
	"fmt"
	"sort"
)

// C использованием библиотеки sort

func main() {
	arr := []int{3, 7, 8, 5, 2, 1, 9, 5, 4}

	sort.Ints(arr)
	// sortedArr := quickSort(arr)
	// fmt.Println("Отсортированный массив:", sortedArr)

	fmt.Println("Отсортированный массив:", arr)
}

// вручную:

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]

	less := []int{}
	equal := []int{}
	greater := []int{}

	for _, num := range arr {
		switch {
		case num < pivot:
			less = append(less, num)
		case num == pivot:
			equal = append(equal, num)
		case num > pivot:
			greater = append(greater, num)
		}
	}

	sortedLess := quickSort(less)
	sortedGreater := quickSort(greater)

	result := append(sortedLess, equal...)
	result = append(result, sortedGreater...)

	return result
}
