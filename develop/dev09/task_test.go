package task

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDownloadSite(t *testing.T) {
	// Создаем тестовый HTTP-сервер
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Отвечаем содержимым для теста
		w.Write([]byte("Test content"))
	}))
	defer testServer.Close()

	// Вызываем downloadSite для тестового URL
	content, err := downloadSite(testServer.URL)
	if err != nil {
		t.Fatalf("Не ожидалась ошибка, получено: %v", err)
	}

	// Проверяем, что содержимое совпадает с ожидаемым
	expectedContent := []byte("Test content")
	if string(content) != string(expectedContent) {
		t.Errorf("Ожидаемый результат: %v, получено: %v", expectedContent, content)
	}
}

func TestMain(m *testing.M) {
	// Запускаем тесты
	exitCode := m.Run()

	// Завершаем выполнение, передавая код выхода
	os.Exit(exitCode)
}
