// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0

package main

import (
	"fmt"
)

func setBit(value int64, position int, bitValue int) int64 {
	// Создаем маску для установки бита в позиции position
	mask := int64(1 << position)

	// Если bitValue равен 1, устанавливаем бит в 1, иначе - в 0
	if bitValue == 1 {
		return value | mask // Устанавливаем бит в 1
	} else {
		return value &^ mask // Устанавливаем бит в 0
	}
}

func main() {
	var value int64 = 15 // Исходное значение (в двоичном формате: 1111)

	// Устанавливаем i-й бит в 1
	i := 2
	value = setBit(value, i, 1)
	fmt.Printf("Значение после установки %d-го бита в 1: %d\n", i, value)

	// Устанавливаем i-й бит в 0
	i = 0
	value = setBit(value, i, 0)
	fmt.Printf("Значение после установки %d-го бита в 0: %d\n", i, value)
}
