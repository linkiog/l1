package main

import (
	"fmt"
	"time"
)

func main() {
	var sec int
	fmt.Println("Введите количество секунд:")
	fmt.Scan(&sec)
	chanData := make(chan int)
	done := make(chan struct{})
	go func() {
		counter := 0
		for {
			select {
			case chanData <- counter:
				counter++
			case <-done:
				close(chanData)
				return

			}
		}
	}()
	go func() {
		for value := range chanData {
			fmt.Printf("Получено значение: %d\n", value)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	time.Sleep(time.Duration(sec * int(time.Second)))
	close(done)
	fmt.Println("Программа завершена.")

}
