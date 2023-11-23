package task

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func grep(lines []string, pattern string, options map[string]interface{}) {
	matches := 0
	contextCount := options["context"].(int)
	afterCount := options["after"].(int)
	beforeCount := options["before"].(int)
	printing := false

	for i, line := range lines {
		match := false

		if options["ignoreCase"].(bool) {
			match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		} else {
			match = strings.Contains(line, pattern)
		}

		if options["invert"].(bool) {
			match = !match
		}

		if match {
			if options["count"].(bool) {
				matches++
				continue
			}

			if options["lineNum"].(bool) {
				fmt.Printf("%d:", i+1)
			}

			// Используем значение флага -C вместо -A и -B, если -C установлен
			if contextCount > 0 {
				// Выводим строки до совпадения
				for j := i - contextCount; j < i; j++ {
					if j >= 0 {
						fmt.Println(lines[j])
					}
				}
				// Выводим строку с совпадением
				fmt.Println(line)
				// Выводим строки после совпадения
				for j := i + 1; j <= i+contextCount && j < len(lines) && j <= i+afterCount; j++ {
					fmt.Println(lines[j])
				}
				printing = true
			} else {
				// Если нет флага -C, используем -B и -A
				for j := i - beforeCount; j < i && j >= 0; j++ {
					fmt.Println(lines[j])
				}

				fmt.Println(line)

				for j := i + 1; j <= i+afterCount && j < len(lines); j++ {
					fmt.Println(lines[j])
				}
				printing = true
			}
		} else if contextCount > 0 {
			// Если была строка с совпадением, выводим оставшийся контекст
			if printing {
				fmt.Println(line)
				contextCount--
			} else if i+contextCount >= len(lines) {
				fmt.Println(line)
			}
		} else {
			printing = false
		}
	}

	if options["count"].(bool) {
		fmt.Println("Matches found:", matches)
	}
}

func main() {
	// Парсим флаги
	before := flag.Int("B", 0, "Print +N lines before a match")
	after := flag.Int("A", 0, "Print +N lines after a match")
	context := flag.Int("C", 0, "Print ±N lines around a match")
	count := flag.Bool("c", false, "Count the number of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case when matching")
	invert := flag.Bool("v", false, "Invert the match, exclude matching lines")
	fixed := flag.Bool("F", false, "Match exact string, not pattern")
	lineNum := flag.Bool("n", false, "Print line numbers")

	flag.Parse()

	// Получаем паттерн из аргументов командной строки
	pattern := flag.Arg(0)

	// Открываем файл для чтения
	fileName := flag.Arg(1)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Читаем строки из файла
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Создаем map опций
	options := map[string]interface{}{
		"before":     *before,
		"after":      *after,
		"context":    *context,
		"count":      *count,
		"ignoreCase": *ignoreCase,
		"invert":     *invert,
		"fixed":      *fixed,
		"lineNum":    *lineNum,
	}

	// Вызываем grep
	grep(lines, pattern, options)
}
