/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.

	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
	через контекст
*/

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Функция генерации данных, которые записываются в канал
func generateData(ctx context.Context, ch chan int) {
	defer close(ch) // Закрыть канал по завершении работы функции
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done(): // Проверить, был ли отменен контекст
			return
		case ch <- i: // Записать данные в канал
		}
	}
	fmt.Println("Нажмите CTRL+C для завершения")
}

// Функция воркера
func worker(ctx context.Context, id int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done() // Уменьшить счетчик группы по завершении работы функции
	for {
		select {
		case <-ctx.Done(): // Проверить, был ли отменен контекст
			return
		case num, ok := <-ch:
			if !ok { // Если канал закрыт, выйти из функции
				return
			}
			fmt.Printf("Worker %d: %d\n", id, num) // Вывести данные в stdout
		}
	}
}

func main() {
	numWorkers := 3 // Количество воркеров
	ch := make(chan int)

	// Создать контекст для отмены в случае получения сигнала завершения
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Отменить контекст при завершении работы программы

	// Запустить горутину для записи данных в канал
	go generateData(ctx, ch)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Запустить горутины для воркеров
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i, &wg, ch)
	}

	// Ожидать сигнала завершения
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel

	// Отменить контекст, чтобы уведомить все горутины о завершении
	cancel()

	// Ожидать завершения работы всех воркеров
	wg.Wait()
}
