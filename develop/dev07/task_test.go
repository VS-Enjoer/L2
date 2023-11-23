package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	// Тест с одним входным каналом, который должен вернуть этот канал
	t.Run("SingleChannel", func(t *testing.T) {
		c := make(chan interface{})
		close(c)

		result := <-or(c)

		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	// Тест с несколькими входными каналами, где первый закрывается
	t.Run("FirstChannelCloses", func(t *testing.T) {
		c1 := make(chan interface{})
		c2 := make(chan interface{})
		close(c1)

		result := <-or(c1, c2)

		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	// Тест с несколькими входными каналами, где второй закрывается
	t.Run("SecondChannelCloses", func(t *testing.T) {
		c1 := make(chan interface{})
		c2 := make(chan interface{})
		close(c2)

		result := <-or(c1, c2)

		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	// Тест с несколькими входными каналами, где первый закрывается позже
	t.Run("FirstChannelClosesLater", func(t *testing.T) {
		c1 := make(chan interface{})
		c2 := make(chan interface{})

		go func() {
			time.Sleep(100 * time.Millisecond)
			close(c1)
		}()

		result := <-or(c1, c2)

		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})
}
