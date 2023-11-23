package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

// cutFields обрабатывает строки, разбивая их на поля и выбирая указанные поля (столбцы) по их номерам.
// Возвращает строку, содержащую результаты выбора полей, разделенных указанным разделителем.
func cutFields(lines []string, fieldNums []int, delimiter string, onlySeparated bool) string {
	var result strings.Builder

	// Итерируемся по каждой строке
	for _, line := range lines {
		// Если требуется выводить только строки содержащие разделитель и текущая строка не содержит разделитель,
		// пропускаем текущую итерацию цикла.
		if onlySeparated && !strings.Contains(line, delimiter) {
			continue
		}

		// Разбиваем строку на слова (поля) с использованием пробела в качестве разделителя.
		fields := strings.Fields(line)

		// Создаем срез для хранения выбранных полей
		var selectedFields []string
		for _, num := range fieldNums {
			// Проверяем что номер поля находится в пределах допустимых значений
			if num > 0 && num <= len(fields) {
				// Добавляем выбранное поле в срез
				selectedFields = append(selectedFields, fields[num-1])
			}
		}

		// Записываем выбранные поля в буфер, разделяя их указанным разделителем
		result.WriteString(strings.Join(selectedFields, delimiter))
		// Добавляем символ новой строки для разделения результатов для каждой строки в исходных данных
		result.WriteString("\n")
	}

	// Возвращаем строку с результатами выбора полей
	return result.String()
}

func main() {
	// Определение флагов командной строки
	var (
		fieldStr      string
		delimiter     string
		onlySeparated bool
	)

	// Регистрация флагов
	flag.StringVar(&fieldStr, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&onlySeparated, "s", false, "только строки с разделителем")
	flag.Parse()

	// Инициализация среза для хранения номеров выбранных полей
	var fieldNums []int
	// Разбиваем строку с номерами полей по запятой и добавляем их в срез
	for _, fieldNum := range strings.Split(fieldStr, ",") {
		if num := strings.TrimSpace(fieldNum); num != "" {
			// Преобразование строки в число и добавление в срез
			fieldNums = append(fieldNums, atoi(num))
		}
	}

	// Чтение строк из стандартного ввода
	lines := readLines()
	// Вызов функции cutFields для обработки строк и вывода результата
	cutFieldsResult := cutFields(lines, fieldNums, delimiter, onlySeparated)
	// Вывод результата
	os.Stdout.WriteString(cutFieldsResult)
}

// atoi преобразует строку в целое число.
func atoi(s string) int {
	num := 0
	for _, digit := range s {
		// Преобразование символа цифры в число и добавление к текущему результату
		num = num*10 + int(digit-'0')
	}
	// Возвращение полученного числа
	return num
}

// readLines считывает строки из стандартного ввода и возвращает их в виде среза строк.
func readLines() []string {
	var lines []string
	// Создание сканнера для чтения из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)
	// Итеративное считывание строк из стандартного ввода
	for scanner.Scan() {
		// Добавление считанной строки в срез строк
		lines = append(lines, scanner.Text())
	}
	// Возвращение среза строк
	return lines
}
