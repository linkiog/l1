package main

import (
	"fmt"
)

// Установить i-й бит в 1
func setBit(num int64, i uint) int64 {
	return num | (1 << i)
}

// Установить i-й бит в 0
func clearBit(num int64, i uint) int64 {
	return num &^ (1 << i)
}

// Проверить значение i-го бита
func checkBit(num int64, i uint) bool {
	return (num & (1 << i)) != 0
}

func main() {
	var num int64 = 42 // Пример числа
	var i uint = 3     // Индекс бита для изменения (счёт от 0)

	fmt.Printf("Изначальное число: %d (%b)\n", num, num)

	// Установить i-й бит в 1
	num = setBit(num, i)
	fmt.Printf("После установки %d-го бита в 1: %d (%b)\n", i, num, num)

	// Установить i-й бит в 0
	num = clearBit(num, i)
	fmt.Printf("После установки %d-го бита в 0: %d (%b)\n", i, num, num)

	// Проверить значение i-го бита
	isSet := checkBit(num, i)
	fmt.Printf("Проверка %d-го бита: %v\n", i, isSet)
}
