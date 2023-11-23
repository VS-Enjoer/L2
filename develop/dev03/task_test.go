package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSortingWithFlags(t *testing.T) {
	tests := []struct {
		inputFile    string
		expectedFile string
		args         []string
	}{
		{"b.txt", "testDirr\\expected\\expected_b.txt", []string{"cmd", "-b"}},
		{"h.txt", "testDirr\\expected\\expected_h.txt", []string{"cmd", "-h"}},
		{"k.txt", "testDirr\\expected\\expected_k.txt", []string{"cmd", "-k", "2"}}, // Добавлена колонка для сортировки
		{"M.txt", "testDirr\\expected\\expected_M.txt", []string{"cmd", "-M"}},
		{"n.txt", "testDirr\\expected\\expected_n.txt", []string{"cmd", "-n"}},
		{"r.txt", "testDirr\\expected\\expected_r.txt", []string{"cmd", "-r"}},
		{"u.txt", "testDirr\\expected\\expected_u.txt", []string{"cmd", "-u"}},
	}

	for _, test := range tests {
		t.Run(test.expectedFile, func(t *testing.T) {
			// Читаем содержимое ожидаемого файла
			expectedContent, err := os.ReadFile(filepath.Join(test.expectedFile))
			if err != nil {
				t.Fatal(err)
			}

			// Создаем временный файл для входных данных
			tmpfile, err := os.CreateTemp("", "testfile*.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			// Пишем входные данные во временный файл
			inputContent, err := os.ReadFile(filepath.Join("testDirr\\input", test.inputFile))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := tmpfile.Write(inputContent); err != nil {
				t.Fatal(err)
			}
			tmpfile.Close()

			// Переопределяем аргументы командной строки
			os.Args = append(test.args, tmpfile.Name())

			// Запускаем программу
			main()

			// Считываем отсортированные строки из временного файла
			sortedContent, err := os.ReadFile("input.txt")
			if err != nil {
				t.Fatal(err)
			}

			// Сравниваем результаты
			if !reflect.DeepEqual(expectedContent, sortedContent) {
				t.Errorf("Ожидалось %s, получено %s", expectedContent, sortedContent)
			}
		})
	}
}
