package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	var sum int
	var mu sync.Mutex
	for _, number := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			sum += n * n
			mu.Unlock()
		}(number)
	}

	wg.Wait()
	fmt.Println(sum)

}
