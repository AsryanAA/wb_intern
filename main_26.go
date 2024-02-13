// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc) функция проверки должна быть регистронезависимой

package main

import (
	"fmt"
	"strings"
)

// Функция проверки уникальности символов в строке
func areAllCharactersUnique(str string) bool {
	// Приводим строку к нижнему регистру
	str = strings.ToLower(str)

	// Создаем карту для подсчета количества встречающихся символов
	charMap := make(map[rune]bool)

	// Проверяем каждый символ в строке
	for _, char := range str {
		// Если символ уже встречался, возвращаем false
		if charMap[char] {
			return false
		}
		// Иначе отмечаем символ как встреченный
		charMap[char] = true
	}

	// Если все символы уникальные, возвращаем true
	return true
}

func main() {
	// Пример строки с уникальными символами
	str1 := "abcdefg"

	// Пример строки с неуникальными символами
	str2 := "Hello"

	// Проверяем уникальность символов в строках
	if areAllCharactersUnique(str1) {
		fmt.Println("Символы в строке str1 уникальны!!!")
	}

	if areAllCharactersUnique(str2) {
		fmt.Println("Символы в строке str2 уникальны!!!")
	}
}
