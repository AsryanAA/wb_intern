/*
	Написать программу, которая конкурентно рассчитает значение
	квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Исходный массив чисел
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем канал для передачи результатов
	resultChan := make(chan int)

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем горутину для каждого числа
	for _, num := range numbers {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			resultChan <- x * x // Вычисляем квадрат числа и отправляем в канал
		}(num)
	}

	// Запускаем горутину для чтения из канала и вывода результата
	go func() {
		wg.Wait()
		close(resultChan) // Закрываем канал после завершения всех горутин
	}()

	// Выводим результаты из канала
	for result := range resultChan {
		fmt.Println(result)
	}
}
