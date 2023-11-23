package main

import (
	"fmt"
	"time"
)

// Функция or принимает один или более done-каналов и возвращает single-канал
func or(channels ...<-chan interface{}) <-chan interface{} {
	// Обрабатываем различные случаи для числа входных каналов
	switch len(channels) {
	case 0:
		// Возвращаем закрытый канал, если нет входных каналов
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		// Возвращаем единственный входной канал, если он есть
		return channels[0]
	}

	// Создаем новый канал для результатов
	orDone := make(chan interface{})

	// Запускаем горутину для выполнения операции выбора
	go func() {
		// Закрываем orDone канал после завершения работы
		defer close(orDone)

		// Используем select с динамическим числом каналов
		select {
		// Если первый канал закрыт, завершаем выбор
		case <-channels[0]:
		// Если второй канал закрыт, завершаем выбор
		case <-channels[1]:
		// Рекурсивно вызываем or для оставшихся каналов
		case <-or(channels[2:]...):
		}
	}()

	// Возвращаем orDone канал
	return orDone
}

func main() {
	// Функция sig создает done-канал, который закроется после указанного времени
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	// Засекаем время начала выполнения
	start := time.Now()

	// Используем функцию or с несколькими done-каналами
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	// Выводим время, прошедшее с начала выполнения
	fmt.Printf("Done after %v", time.Since(start))
}
