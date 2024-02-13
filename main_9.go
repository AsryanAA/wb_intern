/*
	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
	во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

package main

import (
	"fmt"
)

func main() {
	// Создаем каналы для обмена данными
	input := make(chan int)
	output := make(chan int)

	// Функция для чтения чисел из массива и отправки в канал input
	go func() {
		numbers := []int{1, 2, 3, 4, 5} // Пример массива чисел
		for _, num := range numbers {
			input <- num // Отправляем число в канал input
		}
		close(input) // Закрываем канал input после отправки всех чисел
	}()

	// Функция для умножения чисел на 2 и отправки результатов в канал output
	go func() {
		for num := range input {
			output <- num * 2 // Умножаем число на 2 и отправляем результат в канал output
		}
		close(output) // Закрываем канал output после отправки всех результатов
	}()

	// Выводим результаты из канала output в stdout
	for result := range output {
		fmt.Println(result) // Выводим результат в stdout
	}
}
