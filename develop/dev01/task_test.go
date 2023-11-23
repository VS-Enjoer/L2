package main

import (
	"github.com/beevik/ntp"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type mockNtp struct {
	mock.Mock
}

func (m *mockNtp) currentTime() (time.Time, error) {
	args := m.Called()
	return args.Get(0).(time.Time), args.Error(1)
}

func TestMainFunction(t *testing.T) {
	// Создаем экземпляр мок-объекта
	ntpMock := new(mockNtp)

	// Подменяем currentTime на метод мок-объекта
	currentTime = ntpMock.currentTime
	defer func() { currentTime = func() (time.Time, error) { return ntp.Time("0.beevik-ntp.pool.ntp.org") } }()

	// Ожидаем, что метод currentTime будет вызван
	ntpMock.On("currentTime").Return(time.Now(), nil)

	main()

	// Проверяем, что ожидаемый метод был вызван
	ntpMock.AssertExpectations(t)
}
