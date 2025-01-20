package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 8, 9}
	result := make(chan int, len(numbers))
	var wg sync.WaitGroup
	for _, number := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			result <- n * n
		}(number)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	for res := range result {
		fmt.Println(res)
	}
}
