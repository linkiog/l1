package main

import "fmt"

func main() {

	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	for _, str := range strings {
		set[str] = struct{}{}
	}
	for key := range set {
		fmt.Println(key)
	}
}

// без map:

// func main() {
// 	strings := []string{"cat", "cat", "dog", "cat", "tree"}

// 	var set []string

// 	for _, str := range strings {
// 		if !contains(set, str) {
// 			set = append(set, str)
// 		}
// 	}

// 	fmt.Println("Множество:", set)
// }

// func contains(slice []string, value string) bool {
// 	for _, v := range slice {
// 		if v == value {
// 			return true
// 		}
// 	}
// 	return false
// }
