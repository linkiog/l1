package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var workers int
	fmt.Println("Выберите количество workers")
	fmt.Scan(&workers)
	dataChan := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go Workers(ctx, i, dataChan, &wg)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				dataChan <- rand.Intn(100)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	fmt.Println("\nПолучен сигнал завершения, останавливаем воркеров...")
	cancel()
	wg.Wait()
	fmt.Println("Все воркеры завершили работу.")

}
func Workers(ctx context.Context, id int, chanData <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d завершает работу\n", id)
			return
		case data := <-chanData:
			fmt.Printf("Worker %d работает:%d\n", id, data)

		}
	}

}
