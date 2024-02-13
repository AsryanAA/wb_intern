// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20
// НЕ УДАЧНАЯ ПОПЫТКА!!!
package main

import "fmt"

func main() {
	// Задаем значения переменных a и b
	a := 2 << 20
	b := 3 << 20

	fmt.Println(a, b)

	// Выполняем операции
	multiplicationResult := a * b
	divisionResult := a / b
	additionResult := a + b
	subtractionResult := a - b

	// Выводим результаты
	fmt.Println("Результат умножения:", multiplicationResult)
	fmt.Println("Результат деления:", divisionResult)
	fmt.Println("Результат сложения:", additionResult)
	fmt.Println("Результат вычитания:", subtractionResult)
}
