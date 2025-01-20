package main

import "fmt"

func main() {
	a := []int{2, 3, 4, 5, 6, 7, 8}
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan struct{})
	go func() {
		for _, number := range a {
			ch1 <- number
		}
		close(ch1)
	}()

	go func() {
		for ch := range ch1 {
			ch2 <- ch * 2
		}
		close(ch2)
	}()
	go func() {
		for result := range ch2 {
			fmt.Println("Result", result)
		}
		close(done)
	}()
	<-done

}
