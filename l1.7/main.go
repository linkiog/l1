package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	data := make(map[int]int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key, value int) {
			defer wg.Done()
			mu.Lock()
			data[key] = value
			mu.Unlock()
		}(i, i*2)
	}
	wg.Wait()
	fmt.Println(data)

}
