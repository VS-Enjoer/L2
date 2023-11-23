package main

import "fmt"

/*
Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от
своего состояния. Извне создаётся впечатление, что изменился класс объекта.
*/

// Context (Контекст) - смартфон
type Smartphone struct {
	state PhoneState
}

// PhoneState (Состояние телефона) - интерфейс состояния
type PhoneState interface {
	handlePowerButton() string
	handleVolumeButton() string
	handleHomeButton() string
}

// UnlockedState Разблокированный телефон - конкретное состояние
type UnlockedState struct{}

func (u *UnlockedState) handlePowerButton() string {
	return "Заблокировать телефон"
}

func (u *UnlockedState) handleVolumeButton() string {
	return "Регулировать громкость"
}

func (u *UnlockedState) handleHomeButton() string {
	return "Вернуться на главный экран"
}

// LockedState Заблокированный телефон - конкретное состояние
type LockedState struct{}

func (l *LockedState) handlePowerButton() string {
	return "Разблокировать телефон"
}

func (l *LockedState) handleVolumeButton() string {
	return "Показать экран блокировки"
}

func (l *LockedState) handleHomeButton() string {
	return "Показать экран блокировки"
}

// DischargedState Разряженный телефон - конкретное состояние
type DischargedState struct{}

func (d *DischargedState) handlePowerButton() string {
	return "Показать экран зарядки"
}

func (d *DischargedState) handleVolumeButton() string {
	return "Нельзя регулировать громкость, телефон разряжен"
}

func (d *DischargedState) handleHomeButton() string {
	return "Показать экран зарядки"
}

// SetState устанавливает текущее состояние телефона
func (s *Smartphone) SetState(state PhoneState) {
	s.state = state
}

// HandlePowerButton обработка кнопки в текущее состояние
func (s *Smartphone) HandlePowerButton() string {
	return s.state.handlePowerButton()
}

// HandleVolumeButton обработка кнопки в текущее состояние
func (s *Smartphone) HandleVolumeButton() string {
	return s.state.handleVolumeButton()
}

// HandleHomeButton обработка кнопки в текущее состояние
func (s *Smartphone) HandleHomeButton() string {
	return s.state.handleHomeButton()
}

func main() {
	// Создаем экземпляр смартфона
	phone := &Smartphone{}

	// Устанавливаем начальное состояние - разблокированный телефон
	phone.SetState(&UnlockedState{})

	// Имитация нажатия кнопок при разблокированном телефоне
	fmt.Println("Действия с разблокированным телефоном:")
	fmt.Println("Кнопка питания:", phone.HandlePowerButton())
	fmt.Println("Кнопка громкости:", phone.HandleVolumeButton())
	fmt.Println("Кнопка Home:", phone.HandleHomeButton())

	// Переключаем состояние на заблокированный телефон
	phone.SetState(&LockedState{})

	// Имитация нажатия кнопок при заблокированном телефоне
	fmt.Println("\nДействия с заблокированным телефоном:")
	fmt.Println("Кнопка питания:", phone.HandlePowerButton())
	fmt.Println("Кнопка громкости:", phone.HandleVolumeButton())
	fmt.Println("Кнопка Home:", phone.HandleHomeButton())

	// Переключаем состояние на разряженный телефон
	phone.SetState(&DischargedState{})

	// Имитация нажатия кнопок при разряженном телефоне
	fmt.Println("\nДействия с разряженным телефоном:")
	fmt.Println("Кнопка питания:", phone.HandlePowerButton())
	fmt.Println("Кнопка громкости:", phone.HandleVolumeButton())
	fmt.Println("Кнопка Home:", phone.HandleHomeButton())
}
