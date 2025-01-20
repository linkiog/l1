package main

import "fmt"

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	intersection := findIntersection(set1, set2)

	fmt.Println("Пересечение множеств:", intersection)
}

func findIntersection(set1, set2 []int) []int {
	setMap := make(map[int]bool)
	var result []int
	for _, number := range set1 {
		setMap[number] = true
	}
	for _, number := range set2 {
		if setMap[number] {
			result = append(result, number)
		}
	}
	return result
}

//если без мапы:

// func main() {
// 	set1 := []int{1, 2, 3, 4, 5}
// 	set2 := []int{3, 4, 5, 6, 7}

// 	intersection := findIntersection(set1, set2)

// 	fmt.Println("Пересечение множеств:", intersection)
// }
// func findIntersection(set1, set2 []int) []int {
// 	var result []int
// 	for _, num1 := range set1 {
// 		for _, num2 := range set2 {
// 			if num1 == num2 {
// 				if !contains(result, num1) {
// 					result = append(result, num1)
// 				}
// 			}

// 		}
// 	}
// 	return result
// }
// func contains(slice []int, num int) bool {
// 	for _, n := range slice {
// 		if n == num {
// 			return true
// 		}
// 	}
// 	return false

// }
