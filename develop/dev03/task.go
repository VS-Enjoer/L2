package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	column := flag.Int("k", 0, "Указание колонки для сортировки (по умолчанию 0)")
	number := flag.Bool("n", false, "Сортировать по числовому значению")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	uniq := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	month := flag.Bool("M", false, "Сортировать по названию месяца")
	ignoreProb := flag.Bool("b", false, "Игнорировать хвостовые пробелы")
	checkSort := flag.Bool("c", false, "Проверять отсортированы ли данные")
	numberSuffix := flag.Bool("h", false, "Сортировать по числовому значению с учетом суффиксов")

	flag.Parse()

	//Открываем файл который прилетел в аргументах
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		os.Exit(1)
	}
	// Закрываем файл в конце работы программы
	defer file.Close()

	// Читаем строки из файла
	scanner := bufio.NewScanner(file)

	// Массив string для вывода
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()

		// Обработка флага -b (игнорирование хвостовых пробелов)
		if *ignoreProb {
			line = strings.TrimRight(line, " \t")
		}

		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		os.Exit(1)
	}

	// Функция сравнения строк в зависимости от параметров сортировки
	compare := func(i, j int) bool {
		s1 := strings.Fields(lines[i])
		s2 := strings.Fields(lines[j])

		// Учитываем указанную колонку для сортировки
		if *column > 0 && *column <= len(s1) && *column <= len(s2) {
			s1 = s1[*column-1:]
			s2 = s2[*column-1:]
		}

		// Обработка флага -M (сортировка по названию месяца)
		if *month {
			date1, err1 := time.Parse("January", s1[0])
			date2, err2 := time.Parse("January", s2[0])

			// Если ошибки парсинга, то считаем строки равными
			if err1 != nil || err2 != nil {
				return false
			}

			// Сравниваем даты
			if date1.Before(date2) {
				return true
			} else if date1.After(date2) {
				return false
			}
		}

		// Обработка флага -h (сортировка по числовому значению с учетом суффиксов)
		if *numberSuffix {
			num1, err1 := strconv.Atoi(strings.TrimSuffix(s1[0], "n"))
			num2, err2 := strconv.Atoi(strings.TrimSuffix(s2[0], "n"))

			// Если ошибки преобразования, то считаем строки равными
			if err1 != nil || err2 != nil {
				return false
			}

			// Сравниваем числа
			if num1 < num2 {
				return true
			} else if num1 > num2 {
				return false
			}
		}

		// Преобразуем строки в числа, если указан ключ -n
		if *number {
			num1, err1 := strconv.Atoi(s1[0])
			num2, err2 := strconv.Atoi(s2[0])

			// Если ошибки преобразования, то считаем строки равными
			if err1 != nil || err2 != nil {
				return false
			}

			// Сравниваем числа
			if num1 < num2 {
				return true
			} else if num1 > num2 {
				return false
			}
		}

		// Сравниваем строки
		return s1[0] < s2[0]
	}

	// Выполняем сортировку
	sort.Slice(lines, func(i, j int) bool {
		result := compare(i, j)
		if *reverse {
			return !result
		}
		return result
	})

	// Обработка флага -c (проверка отсортированы ли данные)
	if *checkSort {
		for i := 1; i < len(lines); i++ {
			if !compare(i-1, i) {
				fmt.Println("Файл не отсортирован.")
				os.Exit(1)
			}
		}
		fmt.Println("Файл отсортирован.")
		return
	}

	// Убираем повторяющиеся строки, если указан ключ -u
	if *uniq {
		var uniqueLines []string
		seen := make(map[string]bool)
		for _, line := range lines {
			if !seen[line] {
				uniqueLines = append(uniqueLines, line)
				seen[line] = true
			}
		}
		lines = uniqueLines
	}

	// Записываем отсортированные строки в новый файл
	outputFileName := "input.txt"
	err = os.WriteFile(outputFileName, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		os.Exit(1)
	}

	fmt.Println("Отсортированные строки сохранены в файл:", outputFileName)
}
