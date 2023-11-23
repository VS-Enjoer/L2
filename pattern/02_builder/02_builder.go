package main

import "fmt"

/*
Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.
*/

// Car - структура, представляющая автомобиль
type Car struct {
	Model        string
	Color        string
	Engine       string
	Transmission string
}

// CarBuilder - интерфейс строителя для автомобиля
type CarBuilder interface {
	SetModel(model string) CarBuilder
	SetColor(color string) CarBuilder
	SetEngine(engine string) CarBuilder
	SetTransmission(transmission string) CarBuilder
	Build() Car
}

// ConcreteCarBuilder - конкретная реализация строителя
type ConcreteCarBuilder struct {
	car Car
}

func NewConcreteCarBuilder() *ConcreteCarBuilder {
	return &ConcreteCarBuilder{car: Car{}}
}

func (b *ConcreteCarBuilder) SetModel(model string) CarBuilder {
	b.car.Model = model
	return b
}

func (b *ConcreteCarBuilder) SetColor(color string) CarBuilder {
	b.car.Color = color
	return b
}

func (b *ConcreteCarBuilder) SetEngine(engine string) CarBuilder {
	b.car.Engine = engine
	return b
}

func (b *ConcreteCarBuilder) SetTransmission(transmission string) CarBuilder {
	b.car.Transmission = transmission
	return b
}

func (b *ConcreteCarBuilder) Build() Car {
	return b.car
}

// Director - директор, управляющий строителем
type Director struct {
	builder CarBuilder
}

func NewDirector(builder CarBuilder) *Director {
	return &Director{builder: builder}
}

// ConstructCar - метод директора для построения автомобиля
func (d *Director) ConstructCar() Car {
	return d.builder.
		SetModel("Sedan").
		SetColor("Blue").
		SetEngine("V6").
		SetTransmission("Automatic").
		Build()
}

func main() {
	// Создаем конкретного строителя
	builder := NewConcreteCarBuilder()

	// Создаем директора, передаем ему строителя
	director := NewDirector(builder)

	// Используем директора для построения автомобиля
	car := director.ConstructCar()

	// Выводим информацию об автомобиле
	fmt.Println("Car Details:")
	fmt.Printf("- Model: %s\n", car.Model)
	fmt.Printf("- Color: %s\n", car.Color)
	fmt.Printf("- Engine: %s\n", car.Engine)
	fmt.Printf("- Transmission: %s\n", car.Transmission)
}
