package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаём большие числа a и b
	a := big.NewInt(1 << 21) // 2^21 = 2097152
	b := big.NewInt(1 << 22) // 2^22 = 4194304

	// Переменные для хранения результатов
	sum := new(big.Int)
	difference := new(big.Int)
	product := new(big.Int)
	quotient := new(big.Int)

	// Сложение
	sum.Add(a, b)
	fmt.Printf("Сумма: %s\n", sum)

	// Вычитание
	difference.Sub(a, b)
	fmt.Printf("Разность: %s\n", difference)

	// Умножение
	product.Mul(a, b)
	fmt.Printf("Произведение: %s\n", product)

	// Деление (результат деления округляется вниз)
	quotient.Div(a, b)
	fmt.Printf("Частное: %s\n", quotient)
}
