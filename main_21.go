// Реализовать паттерн «адаптер» на любом примере

/*
	Паттерн "Адаптер" (Adapter) позволяет объектам с несовместимыми интерфейсами работать вместе.
	Рассмотрим пример, где у нас есть два различных интерфейса,
	и мы хотим использовать объект, реализующий один интерфейс, с объектом, который ожидает другой интерфейс.
	В данном примере представим, что у нас есть структура Square, которая представляет квадрат, а также интерфейс ShapePrinter,
	который ожидает метод PrintShape для печати формы.
	Однако Square имеет метод Draw для рисования квадрата, который не совместим с PrintShape.
	Мы можем создать адаптер, который реализует интерфейс ShapePrinter и делегирует вызовы метода PrintShape к методу Draw структуры Square.
*/

package main

import "fmt"

// Square представляет квадрат
type Square struct{}

// Draw рисует квадрат
func (s *Square) Draw() {
	fmt.Println("Рисуем квадрат")
}

// ShapePrinter определяет интерфейс для печати формы
type ShapePrinter interface {
	PrintShape()
}

// SquareAdapter адаптер для печати квадрата
type SquareAdapter struct {
	Square *Square
}

// PrintShape печатает квадрат
func (sa *SquareAdapter) PrintShape() {
	sa.Square.Draw()
}

func main() {
	// Создаем квадрат
	square := &Square{}

	// Создаем адаптер для квадрата
	adapter := &SquareAdapter{Square: square}

	// Используем адаптер для печати квадрата
	printShape(adapter)
}

// printShape печатает форму с помощью ShapePrinter
func printShape(sp ShapePrinter) {
	sp.PrintShape()
}
