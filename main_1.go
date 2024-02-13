/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской
	структуры Human (аналог наследования).
*/

package main

import (
	"fmt"
)

// Определяем структуру Human
type Human struct {
	Name string
	Age  int
}

// Метод Speak для структуры Human
func (h Human) Speak() {
	fmt.Printf("Hi, my name is %s and I am %d years old.\n", h.Name, h.Age)
}

// Определяем структуру Action и встраиваем в нее Human
type Action struct {
	Human
}

func main() {
	// Создаем объект типа Action
	action := Action{
		Human: Human{
			Name: "Aram",
			Age:  29,
		},
	}

	// Вызываем метод Speak, который встроен в структуру Action, но наследован из Human
	action.Speak()
}
