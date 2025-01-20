package main

import (
	"fmt"
)

func main() {
	var a interface{} = 42
	var b interface{} = "Hello"
	var c interface{} = true
	var d interface{} = make(chan int)

	checkType(a)
	checkType(b)
	checkType(c)
	checkType(d)
}

func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Тип: int, Значение:", v)
	case string:
		fmt.Println("Тип: string, Значение:", v)
	case bool:
		fmt.Println("Тип: bool, Значение:", v)
	case chan int:
		fmt.Println("Тип: channel, Значение:", v)
	default:
		fmt.Println("Неизвестный тип")
	}
}
