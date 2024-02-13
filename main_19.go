// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode

package main

import (
	"fmt"
)

// Функция для переворачивания строки
func reverseString(str string) string {
	// Преобразуем строку в слайс руны для корректной работы с символами Unicode
	runes := []rune(str)
	length := len(runes)

	// Переворачиваем слайс рун
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	// Возвращаем перевернутую строку
	return string(runes)
}

func main() {
	// Исходная строка
	str := "главрыба — абырвалг"

	// Переворачиваем строку
	reversedStr := reverseString(str)

	// Выводим результат
	fmt.Println("Перевернутая строка:", reversedStr)
}
