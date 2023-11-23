package task

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestExecuteCommand(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "testshell")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		input    string
		expected string
	}{
		{"echo Hello, World!", "Hello, World!\n"},
		{"pwd", getCurrentWorkingDirectory() + "\n"},
		{"echo Test 123", "Test 123\n"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			// Перехватываем стандартный вывод
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Вызываем функцию
			executeCommand(test.input)

			// Восстанавливаем стандартный вывод
			w.Close()
			os.Stdout = oldStdout

			// Сравниваем ожидаемый вывод с фактическим
			var buf bytes.Buffer
			io.Copy(&buf, r)
			actual := buf.String()

			if actual != test.expected {
				t.Errorf("Expected: %s, Got: %s", test.expected, actual)
			}
		})
	}
}

func getCurrentWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir
}
