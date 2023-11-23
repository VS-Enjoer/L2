package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

// Функция, возвращающая текущее время
var currentTime = func() (time.Time, error) {
	return ntp.Time("0.beevik-ntp.pool.ntp.org")
}

func main() {
	// Узнаем точное время c помощью библиотеки NTP
	Time, err := currentTime()
	if err != nil {
		//Если возникает ошибка отправляем ее в stderr и выводим код выхода 1
		fmt.Fprintln(os.Stderr, "Возникла ошибка при получении времени: ", err)
		os.Exit(1)
	}
	fmt.Println("Время по Ntp:", Time.Format(time.RFC3339))
}
