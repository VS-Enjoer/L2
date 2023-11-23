package main

import "fmt"

/*
Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты, позволяя передавать их как
аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.
*/

// Command - интерфейс команды
type Command interface {
	Execute()
}

// Light - получатель (устройство), которое будет управляться
type Light struct{}

func (l *Light) TurnOn() {
	fmt.Println("Light is ON")
}

func (l *Light) TurnOff() {
	fmt.Println("Light is OFF")
}

// LightOnCommand - конкретная реализация команды для включения света
type LightOnCommand struct {
	Light *Light
}

func (c *LightOnCommand) Execute() {
	c.Light.TurnOn()
}

// LightOffCommand - конкретная реализация команды для выключения света
type LightOffCommand struct {
	Light *Light
}

func (c *LightOffCommand) Execute() {
	c.Light.TurnOff()
}

// RemoteControl - инициатор (пульт управления)
type RemoteControl struct {
	Command Command
}

func (r *RemoteControl) PressButton() {
	r.Command.Execute()
}

func main() {
	// Создаем свет и команды для управления им
	light := &Light{}
	lightOnCommand := &LightOnCommand{Light: light}
	lightOffCommand := &LightOffCommand{Light: light}

	// Создаем пульт управления и настраиваем кнопки
	remote := &RemoteControl{}

	// Программируем пульт включить свет
	remote.Command = lightOnCommand
	// Нажимаем кнопку
	remote.PressButton()

	// Программируем пульт выключить свет
	remote.Command = lightOffCommand
	// Нажимаем кнопку
	remote.PressButton()
}
