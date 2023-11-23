package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

// toJSON функция для сериализации объектов в JSON.
func toJSON(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

// parseCreateUpdateParams функция для парсинга и валидации параметров /create_event и /update_event.
func parseCreateUpdateParams(r *http.Request) (Event, error) {
	userIDStr := r.FormValue("user_id")
	date := r.FormValue("date")

	if userIDStr == "" || date == "" {
		return Event{}, errors.New("Неверные параметры запроса")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return Event{}, errors.New("Неверный формат user_id")
	}

	return Event{UserID: userID, Date: date}, nil
}

// jsonResponse функция для отправки JSON-ответа.
func jsonResponse(w http.ResponseWriter, data interface{}) {
	jsonResponse, err := toJSON(data)
	if err != nil {
		http.Error(w, `{"error": "Ошибка сериализации JSON"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// createEventHandler обработчик для метода /create_event.
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	params, err := parseCreateUpdateParams(r)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	err = createEvent(params)
	if err != nil {
		http.Error(w, `{"error": "Ошибка бизнес-логики"}`, http.StatusServiceUnavailable)
		return
	}

	response := map[string]string{"result": "Событие успешно создано"}
	jsonResponse(w, response)
}

// updateEventHandler обработчик для метода /update_event.
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	eventIDStr := r.FormValue("event_id")

	if eventIDStr == "" {
		http.Error(w, `{"error": "Неверные параметры запроса"}`, http.StatusBadRequest)
		return
	}

	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		http.Error(w, `{"error": "Неверный формат event_id"}`, http.StatusBadRequest)
		return
	}

	// Создаем объект Event только с ID и передаем его в updateEvent
	eventToUpdate := Event{ID: eventID}

	err = updateEvent(eventToUpdate)
	if err != nil {
		http.Error(w, `{"error": "Ошибка бизнес-логики"}`, http.StatusServiceUnavailable)
		return
	}

	response := map[string]string{"result": "Событие успешно обновлено"}
	jsonResponse(w, response)
}

// deleteEventHandler обработчик для метода /delete_event.
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	// Распаковка параметров запроса.
	eventID := r.FormValue("event_id")

	// Парсинг параметров запроса.
	if eventID == "" {
		http.Error(w, `{"error": "Неверные параметры запроса"}`, http.StatusBadRequest)
		return
	}

	// Преобразование eventID в int.
	eventIDInt, err := strconv.Atoi(eventID)
	if err != nil {
		http.Error(w, `{"error": "Неверный формат event_id"}`, http.StatusBadRequest)
		return
	}

	// Удаление события.
	err = deleteEvent(eventIDInt)

	// Обработка результата и формирование ответа.
	if err != nil {
		http.Error(w, `{"error": "Ошибка бизнес-логики"}`, http.StatusServiceUnavailable)
		return
	}

	response := map[string]string{"result": "Событие успешно удалено"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// eventsForDayHandler обработчик для метода /events_for_day.
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	// Распаковка параметров запроса.
	userIDStr := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	// Проверка параметров запроса
	if userIDStr == "" || date == "" {
		http.Error(w, `{"error": "Неверные параметры запроса"}`, http.StatusBadRequest)
		return
	}

	// Конверсия userID из строки в int.
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, `{"error": "Неверный формат user_id"}`, http.StatusBadRequest)
		return
	}

	// Получение событий для указанного дня.
	events, err := getEventsForDay(userID, date)

	// Обработка результата и формирование ответа.
	if err != nil {
		http.Error(w, `{"error": "Ошибка бизнес-логики"}`, http.StatusServiceUnavailable)
		return
	}

	// Отправка событий в ответе.
	response := map[string]interface{}{"result": events}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// eventsForWeekHandler обработчик для метода /events_for_week.
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	// Распаковка параметров запроса.
	userIDStr := r.URL.Query().Get("user_id")
	startDate := r.URL.Query().Get("start_date")

	// Парсинг параметров запроса.
	if userIDStr == "" || startDate == "" {
		http.Error(w, `{"error": "Неверные параметры запроса"}`, http.StatusBadRequest)
		return
	}

	// Конвертация userID из строки в int.
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, `{"error": "Неверный формат user_id"}`, http.StatusBadRequest)
		return
	}

	// Получение событий на неделю.
	events, err := getEventsForWeek(userID, startDate)

	// Обработка результата и формирование ответа.
	if err != nil {
		http.Error(w, `{"error": "Ошибка бизнес-логики"}`, http.StatusServiceUnavailable)
		return
	}

	// Отправка событий в ответе.
	response := map[string]interface{}{"result": events}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// eventsForMonthHandler обработчик для метода /events_for_month.
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	// Распаковка параметров запроса.
	userIDStr := r.URL.Query().Get("user_id")
	month := r.URL.Query().Get("month")

	// Парсинг параметров запроса.
	if userIDStr == "" || month == "" {
		http.Error(w, `{"error": "Неверные параметры запроса"}`, http.StatusBadRequest)
		return
	}

	// Конвертация userID из строки в int.
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, `{"error": "Неверный формат user_id"}`, http.StatusBadRequest)
		return
	}

	// Получение событий на месяц.
	events, err := getEventsForMonth(userID, month)

	// Обработка результата и формирование ответа.
	if err != nil {
		http.Error(w, `{"error": "Ошибка бизнес-логики"}`, http.StatusInternalServerError)
		return
	}

	// Отправка событий в ответе.
	response := map[string]interface{}{"result": events}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

// loggingMiddleware промежуточное ПО для логирования запросов.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
