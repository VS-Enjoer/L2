package main

import (
	"errors"
	"time"
)

var events []Event

// Event структура представляет объект доменной области.
type Event struct {
	ID     int    `json:"ID"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
}

// Реализация сохранения события в хранилище.
func createEvent(event Event) error {

	event.ID = len(events) + 1
	events = append(events, event)
	return nil
}

// Реализация получения событий из хранилища.
func getEventsForDay(userID int, date string) ([]Event, error) {

	var result []Event
	for _, event := range events {
		if event.UserID == userID && event.Date == date {
			result = append(result, event)
		}
	}
	return result, nil
}

// Реализация изменения события в хранилища.
func updateEvent(event Event) error {
	for i, existingEvent := range events {
		if existingEvent.ID == event.ID {
			events[i] = event
			return nil
		}
	}
	return errors.New("Событие не найдено")
}

// Реализация удаления события из хранилища.
func deleteEvent(eventID int) error {
	for i, existingEvent := range events {
		if existingEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			return nil
		}
	}
	return errors.New("Событие не найдено")
}

// Реализация получения событий из хранилища(неделя)
func getEventsForWeek(userID int, startDate string) ([]Event, error) {
	var result []Event
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, errors.New("Неверный формат даты")
	}

	for _, event := range events {
		eventDate, err := time.Parse("2006-01-02", event.Date)
		if err == nil && event.UserID == userID && start.Before(eventDate) && eventDate.Before(start.AddDate(0, 0, 7)) {
			result = append(result, event)
		}
	}
	return result, nil
}

// Реализация получения событий из хранилища(месяц)
func getEventsForMonth(userID int, month string) ([]Event, error) {
	var result []Event
	start, err := time.Parse("2006-01", month)
	if err != nil {
		return nil, errors.New("Неверный формат месяца")
	}

	for _, event := range events {
		eventDate, err := time.Parse("2006-01-02", event.Date)
		if err == nil && event.UserID == userID && start.Before(eventDate) && eventDate.Before(start.AddDate(0, 1, 0)) {
			result = append(result, event)
		}
	}
	return result, nil
}
