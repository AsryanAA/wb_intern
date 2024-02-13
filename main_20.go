// Разработать программу, которая переворачивает слова в строке Пример: «snow dog sun — sun dog snow»

package main

import (
	"fmt"
	"strings"
)

// Функция для переворачивания слов в строке
func reverseWords(sentence string) string {
	// Разделяем строку на слова
	words := strings.Fields(sentence)

	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объединяем слова обратно в строку
	return strings.Join(words, " ")
}

func main() {
	// Исходная строка
	// sentence := "snow dog sun"
	sentence := "1 2 3 4 5"

	// Переворачиваем слова в строке
	reversedSentence := reverseWords(sentence)

	// Выводим результат
	fmt.Println("Перевернутые слова:", reversedSentence)
}
