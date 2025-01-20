package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена (context)")
				return
			default:
				fmt.Println("Работаю...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	wg.Wait()
	fmt.Println("Программа завершена.")
}

//channel

// func main() {
// 	stop := make(chan struct{})
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-stop:
// 				fmt.Println("Горутина завершена (канал)")
// 				return
// 			default:
// 				fmt.Println("Работаю...")
// 				time.Sleep(500 * time.Millisecond)
// 			}
// 		}
// 	}()

// 	time.Sleep(2 * time.Second)
// 	close(stop)
// 	wg.Wait()
// 	fmt.Println("Программа завершена")
// }
