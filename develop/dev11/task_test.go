package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateEventHandler(t *testing.T) {
	// Создание фейкового хранилища событий
	events = []Event{}

	// Создание фейкового HTTP запроса
	payload := strings.NewReader("user_id=1&date=2023-01-01")
	req := httptest.NewRequest("POST", "/create_event", payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Создание фейкового HTTP ответа
	w := httptest.NewRecorder()

	// Вызов обработчика
	createEventHandler(w, req)

	// Проверка статус кода
	if w.Code != http.StatusOK {
		t.Errorf("Ожидался код статуса %d, но получен %d", http.StatusOK, w.Code)
	}

	// Проверка наличия события в хранилище
	if len(events) != 1 {
		t.Errorf("Ожидалось, что событие будет добавлено в хранилище")
	}

	// Проверка содержимого ответа
	expectedResponse := `{"result":"Событие успешно создано"}`
	if strings.TrimSpace(w.Body.String()) != expectedResponse {
		t.Errorf("Ожидался ответ %s, но получен %s", expectedResponse, w.Body.String())
	}
}

func TestDeleteEventHandler(t *testing.T) {
	// Создание фейкового хранилища событий
	events = []Event{{ID: 1, UserID: 1, Date: "2023-01-01"}}

	// Создание фейкового HTTP запроса
	req := httptest.NewRequest("DELETE", "/delete_event?event_id=1", nil)

	// Создание фейкового HTTP ответа
	w := httptest.NewRecorder()

	// Вызов обработчика
	deleteEventHandler(w, req)

	// Проверка статус кода
	if w.Code != http.StatusOK {
		t.Errorf("Ожидался код статуса %d, но получен %d", http.StatusOK, w.Code)
	}

	// Проверка удаления события из хранилища
	if len(events) != 0 {
		t.Errorf("Ожидалось, что событие будет удалено из хранилища")
	}

	// Проверка содержимого ответа
	expectedResponse := `{"result":"Событие успешно удалено"}`
	if strings.TrimSpace(w.Body.String()) != expectedResponse {
		t.Errorf("Ожидался ответ %s, но получен %s", expectedResponse, w.Body.String())
	}
}

func TestEventsForDayHandler(t *testing.T) {
	// Создание фейкового хранилища событий
	events = []Event{{ID: 1, UserID: 1, Date: "2023-01-01"}}

	// Создание фейкового HTTP запроса
	req := httptest.NewRequest("GET", "/events_for_day?user_id=1&date=2023-01-01", nil)

	// Создание фейкового HTTP ответа
	w := httptest.NewRecorder()

	// Вызов обработчика
	eventsForDayHandler(w, req)

	// Проверка статус кода
	if w.Code != http.StatusOK {
		t.Errorf("Ожидался код статуса %d, но получен %d", http.StatusOK, w.Code)
	}

	// Проверка содержимого ответа
	expectedResponse := `{"result":[{"ID":1,"user_id":1,"date":"2023-01-01"}]}`
	if strings.TrimSpace(w.Body.String()) != expectedResponse {
		t.Errorf("Ожидался ответ %s, но получен %s", expectedResponse, w.Body.String())
	}
}

//
//func TestUpdateEventHandler(t *testing.T) {
//	// Создание фейкового хранилища событий
//	events = []Event{{ID: 1, UserID: 1, Date: "2023-01-01"}}
//
//	// Создание фейкового HTTP запроса
//	payload := strings.NewReader("user_id=1&date=2023-01-02")
//	req := httptest.NewRequest("POST", "/update_event", payload)
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//
//	// Предположим, что ID события 1
//	req.Form = map[string][]string{
//		"event_id": {"1"},
//	}
//
//	// Создание фейкового HTTP ответа
//	w := httptest.NewRecorder()
//
//	// Вызов обработчика
//	updateEventHandler(w, req)
//
//	// Проверка статус кода
//	if w.Code != http.StatusOK {
//		t.Errorf("Ожидался код статуса %d, но получен %d. Тело ответа: %s", http.StatusOK, w.Code, w.Body.String())
//	}
//
//	// Проверка обновления события в хранилище
//	updatedEvent := events[0]
//	if updatedEvent.Date != "2023-01-02" {
//		t.Errorf("Ожидалось, что cобытие успешно обновлено")
//	}
//
//	// Проверка содержимого ответа
//	expectedResponse := `{"result":"Событие успешно обновлено"}`
//	if strings.TrimSpace(w.Body.String()) != expectedResponse {
//		t.Errorf("Ожидался ответ %s, но получен %s", expectedResponse, w.Body.String())
//	}
//}
