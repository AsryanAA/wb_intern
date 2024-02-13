// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика
// через атомик

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Структура для счетчика
type Counter struct {
	value int64 // Используем int64 для атомической операции
}

// Метод для инкрементации счетчика
func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Метод для получения текущего значения счетчика
func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	// Создаем новый экземпляр счетчика
	counter := Counter{}

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем несколько горутин для инкрементации счетчика
	for i := 0; i < 42; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}
