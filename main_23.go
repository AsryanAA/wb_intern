// Удалить i-ый элемент из слайса

package main

import "fmt"

func main() {
	// Исходный срез
	slice := []int{1, 2, 42, 4, 5}

	// Индекс элемента, который хотим удалить
	index := 2

	// Удаляем элемент
	slice = append(slice[:index], slice[index+1:]...)

	// Выводим результат
	fmt.Println("Срез после удаления элемента:", slice)
}
