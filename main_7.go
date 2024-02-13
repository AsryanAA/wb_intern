// Реализовать конкурентную запись данных в map

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем мьютекс для безопасного доступа к map
	var mu sync.Mutex
	// Создаем map для хранения данных
	data := make(map[string]int)

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Определяем функцию для конкурентной записи данных в map
	writeToMap := func(key string, value int) {
		// Захватываем мьютекс перед записью в map
		mu.Lock()
		defer mu.Unlock()
		// Записываем данные в map
		data[key] = value
	}

	// Запускаем горутины для конкурентной записи данных в map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d->", index)
			value := index * 10
			writeToMap(key, value)
		}(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим содержимое map после всех операций записи
	fmt.Println("Map data", data)
}
