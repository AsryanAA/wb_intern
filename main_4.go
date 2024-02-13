/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.

	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
	через канал
*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Функция генерации данных, которые записываются в канал
func generateData(ch chan int) {
	defer close(ch) // Закрыть канал по завершении работы функции
	for i := 1; i <= 10; i++ {
		ch <- i // Записать данные в канал
	}
	fmt.Println("Нажмите CTRL+C для завершения")
}

// Функция воркера
func worker(id int, wg *sync.WaitGroup, ch chan int, done chan struct{}) {
	defer wg.Done() // Уменьшить счетчик группы по завершении работы функции
	for {
		select {
		case num, ok := <-ch:
			if !ok { // Если канал закрыт, выйти из функции
				return
			}
			fmt.Printf("Worker %d: %d\n", id, num) // Вывести данные в stdout
		case <-done: // Проверить сигнал завершения
			return
		}
	}
}

func main() {
	numWorkers := 3 // Количество воркеров
	ch := make(chan int)
	done := make(chan struct{})

	// Запустить горутину для записи данных в канал
	go generateData(ch)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Запустить горутины для воркеров
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg, ch, done)
	}

	// Ожидать сигнала завершения

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel

	// Закрыть канал завершения, чтобы уведомить воркеров о завершении
	close(done)

	// Ожидать завершения работы всех воркеров
	wg.Wait()
}
