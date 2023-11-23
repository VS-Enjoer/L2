package main

import "fmt"

/*
Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по
цепочке обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос
дальше по цепи.
*/

// Request - структура, представляющая запрос
type Request struct {
	Path            string
	UserID          int
	IsAuthenticated bool
}

// Handler - интерфейс обработчика запроса
type Handler interface {
	HandleRequest(request *Request) // Метод для обработки запроса
	SetNext(handler Handler)        // Метод для установки следующего обработчика в цепи
}

// AuthenticationHandler - обработчик для проверки аутентификации
type AuthenticationHandler struct {
	next Handler
}

func (h *AuthenticationHandler) HandleRequest(request *Request) {
	// Проверяем, прошла ли аутентификация
	if !request.IsAuthenticated {
		fmt.Println("Ошибка аутентификации.")
		return
	}
	fmt.Println("Прошли аутентификацию.")
	// Если есть следующий обработчик, передаем ему запрос
	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

func (h *AuthenticationHandler) SetNext(handler Handler) {
	h.next = handler
}

// AuthorizationHandler - обработчик для проверки прав доступа
type AuthorizationHandler struct {
	next Handler
}

func (h *AuthorizationHandler) HandleRequest(request *Request) {
	// Проверяем, есть ли права доступа
	if request.UserID != 123 {
		fmt.Println("Ошибка авторизации.")
		return
	}
	fmt.Println("Успешная авторизация.")
	// Если есть следующий обработчик, передаем ему запрос
	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

func (h *AuthorizationHandler) SetNext(handler Handler) {
	h.next = handler
}

// LoggingHandler - обработчик для логирования
type LoggingHandler struct {
	next Handler
}

func (h *LoggingHandler) HandleRequest(request *Request) {
	// Логируем информацию о запросе
	fmt.Printf("Logging: %s\n", request.Path)
	// Если есть следующий обработчик, передаем ему запрос
	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

func (h *LoggingHandler) SetNext(handler Handler) {
	h.next = handler
}

func main() {
	// Создаем обработчики запроса
	authenticationHandler := &AuthenticationHandler{}
	authorizationHandler := &AuthorizationHandler{}
	loggingHandler := &LoggingHandler{}

	// Связываем обработчики в цепочку
	authenticationHandler.SetNext(authorizationHandler)
	authorizationHandler.SetNext(loggingHandler)

	// Создаем запрос
	request := &Request{
		Path:            "/secure-page",
		UserID:          123,
		IsAuthenticated: true,
	}

	// Обработка запроса, который проходит через цепочку обработчиков
	authenticationHandler.HandleRequest(request)
}
