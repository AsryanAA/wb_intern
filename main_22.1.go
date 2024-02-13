package main

import (
	"fmt"
	"math/big"
	"math"
)

func main() {
	// Умножение
	num1 := big.NewInt(int64(math.Pow(2, 20)))
	num2 := big.NewInt(int64(math.Pow(3, 20)))
	result := new(big.Int).Mul(num1, num2)
	fmt.Println("Умножение:", result)

	// Деление
	dividend := big.NewInt(int64(math.Pow(2, 20)))
	divisor := big.NewInt(int64(math.Pow(3, 20)))
	quotient := new(big.Int).Div(dividend, divisor)
	fmt.Println("Деление:", quotient)

	// Сложение
	summand1 := big.NewInt(int64(math.Pow(2, 20)))
	summand2 := big.NewInt(int64(math.Pow(3, 20)))
	sum := new(big.Int).Add(summand1, summand2)
	fmt.Println("Сложение:", sum)

	// Вычитание
	minuend := big.NewInt(int64(math.Pow(2, 20)))
	subtrahend := big.NewInt(int64(math.Pow(3, 20)))
	difference := new(big.Int).Sub(minuend, subtrahend)
	fmt.Println("Вычитание:", difference)
}