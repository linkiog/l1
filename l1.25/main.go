package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) >= duration {
			break
		}
	}
}

func main() {
	fmt.Println("Начало работы")
	Sleep(2 * time.Second)
	fmt.Println("Завершение работы")
}
