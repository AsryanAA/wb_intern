// Реализовать бинарный поиск встроенными методами языка

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Отсортированный массив
	arr2 := []int{2, 5, 7, 9, 13, 16, 21, 26, 31}
	arr := []int{31, 26, 21, 16, 13, 9, 7, 5, 2}
	arr1 := []int{10, 18, -2, 3, 56, 42, 1, 10, 2}
	sort.Ints(arr1)
	fmt.Println(arr2)

	// Элемент, который мы ищем
	target := 31

	// Бинарный поиск элемента в массиве
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] <= target // знак >= или <= зависит от того в какой последовательности отсортирован массив (убывающая)
	})
	fmt.Println(index)
	// проверка не обязательна
	// Проверяем, был ли найден элемент и выводим результат
	if index < len(arr) && arr[index] == target {
		fmt.Printf("Элемент %d найден в позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
