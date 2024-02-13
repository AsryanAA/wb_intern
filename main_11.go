// Реализовать пересечение двух неупорядоченных множеств

package main

import "fmt"

// Функция для нахождения пересечения двух множеств
func intersection(set1, set2 []int) []int {
	// Создаем карту для хранения элементов первого множества
	set1Map := make(map[int]bool)

	// Заполняем карту элементами первого множества
	for _, num := range set1 {
		set1Map[num] = true
	}

	// Создаем слайс для хранения результатов пересечения
	var intersect []int

	// Проверяем элементы второго множества и добавляем их в результат, если они есть в первом множестве
	for _, num := range set2 {
		if set1Map[num] {
			intersect = append(intersect, num)
		}
	}

	return intersect
}

func main() {
	// Пример неупорядоченных множеств
	set1 := []int{1, 3, 5, 7, 9}
	set2 := []int{2, 3, 5, 8, 9}

	// Находим пересечение множеств
	intersect := intersection(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение множеств:", intersect)
}
