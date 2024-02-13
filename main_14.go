// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}

package main

import (
	"fmt"
)

func main() {
	var val interface{}
	var valInt8 int8

	val = 42
	checkType(val)

	valInt8 = 42
	checkType(valInt8) // не известный тип переменной

	val = "Hello"
	checkType(val)

	val = true
	checkType(val)

	val = make(chan int)
	checkType(val)

	val = 1.5
	checkType(val) // не известный тип переменной
}

func checkType(val interface{}) {
	// Используем типовое утверждение для проверки типа переменной
	switch val.(type) {
	case int:
		fmt.Println("Тип переменной: целое число(int)")
	case string:
		fmt.Println("Тип переменной: строка(string)")
	case bool:
		fmt.Println("Тип переменной: логическое true or false (bool)")
	case chan int:
		fmt.Println("Тип переменной: канал (channel)")
	default:
		fmt.Println("Неизвестный тип переменной")
	}
}
