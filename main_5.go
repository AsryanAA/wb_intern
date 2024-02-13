package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 1; ; i++ {
		ch <- i // Отправляем значение в канал
	}
}

func main() {
	duration := 3 * time.Second // Время работы программы (5 секунд)
	ch := make(chan int)

	// Запускаем горутину для отправки значений в канал
	go sender(ch)

	// Читаем значения из канала в основной горутине
	start := time.Now()
	for {
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
		case <-time.After(duration - time.Since(start)):
			fmt.Println("Time's up! Exiting...")
			return
		}
	}
}
