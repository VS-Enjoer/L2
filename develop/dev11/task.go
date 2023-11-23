package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8080"

	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", eventsForDayHandler)
	http.HandleFunc("/events_for_week", eventsForWeekHandler)
	http.HandleFunc("/events_for_month", eventsForMonthHandler)

	// Добавление middleware для логирования.
	http.Handle("/", loggingMiddleware(http.DefaultServeMux))

	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
