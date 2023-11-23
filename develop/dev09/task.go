package task

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Парсим аргументы командной строки
	urlFlag := flag.String("url", "", "URL сайта для скачивания")
	flag.Parse()

	// Проверяем, был ли передан URL
	if *urlFlag == "" {
		fmt.Println("Использование: go run main.go -url <URL>")
		os.Exit(1)
	}

	// Получаем содержимое страницы
	content, err := downloadSite(*urlFlag)
	if err != nil {
		fmt.Printf("Ошибка при скачивании сайта: %v\n", err)
		os.Exit(1)
	}

	// Создаем каталог для сохранения содержимого
	baseURL, err := url.Parse(*urlFlag)
	if err != nil {
		fmt.Printf("Ошибка при анализе URL: %v\n", err)
		os.Exit(1)
	}
	dirName := strings.Replace(baseURL.Host, ".", "_", -1)
	err = os.Mkdir(dirName, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Ошибка при создании каталога: %v\n", err)
		os.Exit(1)
	}

	// Сохраняем содержимое в файл
	filePath := filepath.Join(dirName, "index.html")
	err = os.WriteFile(filePath, content, os.ModePerm)
	if err != nil {
		fmt.Printf("Ошибка при записи в файл: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Сайт успешно скачан в %s\n", filePath)
}

func downloadSite(url string) ([]byte, error) {
	// Выполняем HTTP-запрос
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Читаем содержимое страницы
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
