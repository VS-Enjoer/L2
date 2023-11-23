package task

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	// Тест совпадения по строке "образец"
	t.Run("MatchPattern", func(t *testing.T) {
		lines := []string{"Первая строка", "Второй образец строки", "Строка бум.", "Гига строка.", "И снова образец", "End строка"}
		pattern := "образец"
		options := map[string]interface{}{"before": 0, "after": 0, "context": 0, "count": false, "ignoreCase": false, "invert": false, "fixed": false, "lineNum": false}
		grepResult := captureOutput(func() {
			grep(lines, pattern, options)
		})
		expectedOutput := "Второй образец строки\nИ снова образец\n"
		if grepResult != expectedOutput {
			t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, grepResult)
		}
	})

	// Тест совпадения с игнорированием регистра
	t.Run("IgnoreCase", func(t *testing.T) {
		lines := []string{"Первая строка", "Второй образец строки", "Строка бум.", "Гига строка.", "И снова образец", "End строка"}
		pattern := "ОбРаЗец"
		options := map[string]interface{}{"before": 0, "after": 0, "context": 0, "count": false, "ignoreCase": true, "invert": false, "fixed": false, "lineNum": false}
		grepResult := captureOutput(func() {
			grep(lines, pattern, options)
		})
		expectedOutput := "Второй образец строки\nИ снова образец\n"
		if grepResult != expectedOutput {
			t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, grepResult)
		}
	})

	// Тест инвертированного совпадения
	t.Run("InvertMatch", func(t *testing.T) {
		lines := []string{"Первая строка", "Второй образец строки", "Строка бум.", "Гига строка.", "И снова образец", "End строка"}
		pattern := "образец"
		options := map[string]interface{}{"before": 0, "after": 0, "context": 0, "count": false, "ignoreCase": false, "invert": true, "fixed": false, "lineNum": false}
		grepResult := captureOutput(func() {
			grep(lines, pattern, options)
		})
		expectedOutput := "Первая строка\nСтрока бум.\nГига строка.\nEnd строка\n"
		if grepResult != expectedOutput {
			t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, grepResult)
		}
	})
}

// Функция для захвата вывода в консоль
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var output strings.Builder
	io.Copy(&output, r)
	return output.String()
}
