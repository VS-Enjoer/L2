package main

import "fmt"

/*
Паттерн проектирования «Посетитель» позволяет открепить алгоритм от структуры того объекта, которым он оперирует.
Практический результат такого открепления – способность добавлять новые операции к имеющимся структурам объектов, не модифицируя эти структуры.
*/

// CarComponent - интерфейс компонента машины
type CarComponent interface {
	Accept(Visitor)
}

// Engine - конкретная реализация компонента (двигатель)
type Engine struct {
	Status string
}

func (e *Engine) Accept(visitor Visitor) {
	visitor.VisitEngine(e)
}

// Wheel - конкретная реализация компонента (колесо)
type Wheel struct {
	Status string
}

func (w *Wheel) Accept(visitor Visitor) {
	visitor.VisitWheel(w)
}

// Body - конкретная реализация компонента (кузов)
type Body struct {
	Status string
}

func (b *Body) Accept(visitor Visitor) {
	visitor.VisitBody(b)
}

// Visitor - интерфейс посетителя
type Visitor interface {
	VisitEngine(engine *Engine)
	VisitWheel(wheel *Wheel)
	VisitBody(body *Body)
}

// TechnicalVisitor - конкретная реализация посетителя для проверки технического состояния
type TechnicalVisitor struct{}

func (t *TechnicalVisitor) VisitEngine(engine *Engine) {
	// Проводим технический осмотр двигателя
	engine.Status = "OK"
}

func (t *TechnicalVisitor) VisitWheel(wheel *Wheel) {
	// Проводим технический осмотр колеса
	wheel.Status = "OK"
}

func (t *TechnicalVisitor) VisitBody(body *Body) {
	// Проводим технический осмотр кузова
	body.Status = "OK"
}

func main() {
	// Создаем компоненты машины
	engine := &Engine{Status: "Нужна проверка"}
	wheel := &Wheel{Status: "Нужна проверка"}
	body := &Body{Status: "Нужна проверка"}

	// Создаем посетителя для технического осмотра
	technicalVisitor := &TechnicalVisitor{}

	// Применяем посетителя к каждому компоненту
	engine.Accept(technicalVisitor)
	wheel.Accept(technicalVisitor)
	body.Accept(technicalVisitor)

	// Выводим результаты технического осмотра
	fmt.Printf("Engine status: %s\n", engine.Status)
	fmt.Printf("Wheel status: %s\n", wheel.Status)
	fmt.Printf("Body status: %s\n", body.Status)
}
