package main

import "fmt"

/*
Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый
из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
*/

// MapStrategy - интерфейс стратегии для прокладывания маршрута
type MapStrategy interface {
	BuildRoute(start, end string) string
}

// Navigator - контекст, который использует стратегию
type Navigator struct {
	strategy MapStrategy
}

// SetStrategy - устанавливает стратегию прокладывания маршрута
func (n *Navigator) SetStrategy(strategy MapStrategy) {
	n.strategy = strategy
}

// PlanRoute - планирует маршрут с использованием текущей стратегии
func (n *Navigator) PlanRoute(start, end string) {
	route := n.strategy.BuildRoute(start, end)
	fmt.Println("Маршрут:", route)
}

// WalkStrategy - конкретная стратегия для прокладывания маршрута пешехода
type WalkStrategy struct{}

func (w *WalkStrategy) BuildRoute(start, end string) string {
	return fmt.Sprintf("Пешеходный маршрут построен из %s в %s", start, end)
}

// BusStrategy - конкретная стратегия для прокладывания маршрута автобуса
type BusStrategy struct{}

func (b *BusStrategy) BuildRoute(start, end string) string {
	return fmt.Sprintf("Автобусный маршрут построен из %s в %s", start, end)
}

// CarStrategy - конкретная стратегия для прокладывания маршрута автомобиля
type CarStrategy struct{}

func (c *CarStrategy) BuildRoute(start, end string) string {
	return fmt.Sprintf("Автомобильный маршрут построен из %s в %s", start, end)
}

func main() {
	// Создаем экземпляр навигатора
	navigator := &Navigator{}

	// Устанавливаем стратегию прокладывания маршрута для пешехода и планируем маршрут
	navigator.SetStrategy(&WalkStrategy{})
	navigator.PlanRoute("Точка А", "Точка B")

	// Устанавливаем стратегию прокладывания маршрута для автобуса и планируем маршрут
	navigator.SetStrategy(&BusStrategy{})
	navigator.PlanRoute("Точка А", "Точка B")

	// Устанавливаем стратегию прокладывания маршрута для автомобиля и планируем маршрут
	navigator.SetStrategy(&CarStrategy{})
	navigator.PlanRoute("Точка А", "Точка B")
}
