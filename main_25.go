// Реализовать собственную функцию sleep

package main

import (
	"fmt"
	"time"
)

// Sleep ожидает указанное количество миллисекунд
func Sleep(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

func main() {
	fmt.Println("Начало ожидания...")
	Sleep(3000) // Ожидаем 3 секунды
	fmt.Println("Прошло", time.Duration(3000)*time.Millisecond)
	fmt.Println("Ожидание завершено!")
}
