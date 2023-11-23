package task

import (
	"bytes"
	"fmt"
	"net"
	"testing"
	"time"
)

// FakeNetworkHandler фейковая реализация NetworkHandler для тестирования
type FakeNetworkHandler struct {
	DialTimeoutFunc func(network, address string, timeout time.Duration) (net.Conn, error)
	ListenFunc      func(network, address string) (net.Listener, error)
	AcceptFunc      func(l net.Listener) (net.Conn, error)
	CloseFunc       func(c net.Conn) error
}

func (f FakeNetworkHandler) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return f.DialTimeoutFunc(network, address, timeout)
}

func (f FakeNetworkHandler) Listen(network, address string) (net.Listener, error) {
	return f.ListenFunc(network, address)
}

func (f FakeNetworkHandler) Accept(l net.Listener) (net.Conn, error) {
	return f.AcceptFunc(l)
}

func (f FakeNetworkHandler) Close(c net.Conn) error {
	return f.CloseFunc(c)
}

// TestConnectAndCopy функция тестирования ConnectAndCopy
func TestConnectAndCopy(t *testing.T) {
	t.Run("TestSuccessfulConnection", func(t *testing.T) {
		// Инициализация FakeNetworkHandler с успешным DialTimeout
		fakeHandler := FakeNetworkHandler{
			DialTimeoutFunc: func(network, address string, timeout time.Duration) (net.Conn, error) {
				return &FakeConn{}, nil
			},
			CloseFunc: func(c net.Conn) error {
				return nil
			},
		}

		// Вызов ConnectAndCopy с успешным подключением
		err := ConnectAndCopy(fakeHandler, "tcp", "example.com:8080", 10*time.Second, bytes.NewReader([]byte("test")), &bytes.Buffer{})

		// Проверка на отсутствие ошибок
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("TestFailedConnection", func(t *testing.T) {
		// Инициализация FakeNetworkHandler с ошибкой при DialTimeout
		fakeHandler := FakeNetworkHandler{
			DialTimeoutFunc: func(network, address string, timeout time.Duration) (net.Conn, error) {
				return nil, fmt.Errorf("connection error")
			},
		}

		// Вызов ConnectAndCopy с ошибкой при подключении
		err := ConnectAndCopy(fakeHandler, "tcp", "example.com:8080", 10*time.Second, bytes.NewReader([]byte("test")), &bytes.Buffer{})

		// Проверка на наличие ошибки
		if err == nil {
			t.Error("Expected an error, got nil")
		}
	})
}

// FakeConn фейковая реализация net.Conn для тестирования
type FakeConn struct {
	bytes.Buffer
}

func (f *FakeConn) Close() error {
	return nil
}

func (f *FakeConn) LocalAddr() net.Addr {
	return &net.IPAddr{IP: net.IPv4(127, 0, 0, 1)}
}

func (f *FakeConn) RemoteAddr() net.Addr {
	return &net.IPAddr{IP: net.IPv4(192, 168, 0, 1)}
}

func (f *FakeConn) SetDeadline(t time.Time) error {
	return nil
}

func (f *FakeConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (f *FakeConn) SetWriteDeadline(t time.Time) error {
	return nil
}
