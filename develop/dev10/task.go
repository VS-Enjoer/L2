package task

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// NetworkHandler интерфейс для действий с сетью
type NetworkHandler interface {
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
	Listen(network, address string) (net.Listener, error)
	Accept(l net.Listener) (net.Conn, error)
	Close(c net.Conn) error
}

// NetworkHandlerImpl реализация NetworkHandler
type NetworkHandlerImpl struct{}

func (nh NetworkHandlerImpl) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout(network, address, timeout)
}

func (nh NetworkHandlerImpl) Listen(network, address string) (net.Listener, error) {
	return net.Listen(network, address)
}

func (nh NetworkHandlerImpl) Accept(l net.Listener) (net.Conn, error) {
	return l.Accept()
}

func (nh NetworkHandlerImpl) Close(c net.Conn) error {
	return c.Close()
}

// ConnectAndCopy подключается к серверу и копирует данные между сокетом и указанными источником и приемником данных
func ConnectAndCopy(handler NetworkHandler, network, address string, timeout time.Duration, source io.Reader, destination io.Writer) error {
	// Формирование адреса для подключения
	conn, err := handler.DialTimeout(network, address, timeout)
	if err != nil {
		return fmt.Errorf("Ошибка подключения: %v", err)
	}
	defer handler.Close(conn)

	// Запуск горутины для чтения данных из сокета и вывода их в destination
	go func() {
		io.Copy(destination, conn)
		fmt.Fprintln(destination, "Сервер закрыл соединение.")
	}()

	// Копирование данных из source в сокет
	_, err = io.Copy(conn, source)
	if err != nil {
		return fmt.Errorf("Ошибка при отправке данных: %v", err)
	}

	return nil
}

func main() {
	// Парсинг аргументов командной строки
	host := flag.String("host", "", "Хост (IP или доменное имя)")
	port := flag.String("port", "", "Порт")
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут подключения")

	flag.Parse()

	// Проверка обязательных параметров
	if *host == "" || *port == "" {
		fmt.Println("Необходимо указать хост и порт.")
		flag.PrintDefaults()
		return
	}

	handler := NetworkHandlerImpl{}

	// Вызов функции ConnectAndCopy
	err := ConnectAndCopy(handler, "tcp", fmt.Sprintf("%s:%s", *host, *port), *timeout, os.Stdin, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
