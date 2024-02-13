// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество

package main

import "fmt"

// Функция для создания множества из последовательности строк
func makeSet(strings []string) map[string]bool {
	set := make(map[string]bool)
	for _, str := range strings {
		set[str] = true // Добавляем строку в множество как ключ карты
	}
	return set
}

func main() {
	// Исходная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество из последовательности строк
	set := makeSet(strings)

	// Выводим множество
	fmt.Println("Множество:", set)
}
