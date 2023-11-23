package main

import (
	"testing"
)

func TestCutFields(t *testing.T) {
	// Тестовые данные
	lines := []string{
		"Петя Вася Ваня",
		"1 2 3",
		"Михал_Палыч_Терентьев Норм Мужик",
	}

	// Ожидаемый результат для выбранных полей 1 и 3 с пробелом в качестве разделителя
	expected := "Петя;Ваня\n1;3\nМихал_Палыч_Терентьев;Мужик\n"

	// Запускаем функцию cutFields
	result := cutFields(lines, []int{1, 3}, ";", false)

	// Сравниваем результат с ожидаемым
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	// Тест с другими полями и разделителем
	expected2 := "Вася,Ваня\n2,3\nНорм,Мужик\n"
	result2 := cutFields(lines, []int{2, 3}, ",", false)
	if result2 != expected2 {
		t.Errorf("Expected: %s, Got: %s", expected2, result2)
	}

	// Тест с выбором только строк, содержащих разделитель
	expected3 := "Петя\n1\nМихал_Палыч_Терентьев\n"
	result3 := cutFields(lines, []int{1}, " ", true)
	if result3 != expected3 {
		t.Errorf("Expected: %s, Got: %s", expected3, result3)
	}
}
